package cluster

import (
	"fmt"

	awsSdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/openshift/osdctl/pkg/osdCloud"
	"github.com/openshift/osdctl/pkg/provider/aws"
	"github.com/openshift/osdctl/pkg/utils"
	"github.com/spf13/cobra"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

type cpdOptions struct {
	clusterID  string
	awsProfile string
}

const (
	unknownProvisionCode = "OCM3999"
	cpdLongDescription   = `
Helps investigate OSD/ROSA cluster provisioning delays (CPD) or failures

  This command only supports AWS at the moment and will:
	
  * Check the cluster's dnszone.hive.openshift.io custom resource
  * Check whether a known OCM error code and message has been shared with the customer already
  * Check that the cluster's VPC and/or subnet route table(s) contain a route for 0.0.0.0/0 if it's BYOVPC
`
	cpdExample = `
  # Investigate a CPD for a cluster using an AWS profile named "rhcontrol"
  osdctl cluster cpd --cluster-id 1kfmyclusteristhebesteverp8m --profile rhcontrol
`
)

func newCmdCpd() *cobra.Command {
	ops := cpdOptions{}
	cpdCmd := &cobra.Command{
		Use:               "cpd",
		Short:             "Runs diagnostic for a Cluster Provisioning Delay (CPD)",
		Long:              cpdLongDescription,
		Example:           cpdExample,
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {

			cmdutil.CheckErr(ops.run())
		},
	}
	cpdCmd.Flags().StringVarP(&ops.clusterID, "cluster-id", "C", ops.clusterID, "The internal (OCM) Cluster ID")
	cpdCmd.Flags().StringVarP(&ops.awsProfile, "profile", "p", ops.awsProfile, "AWS profile name")

	return cpdCmd
}

func (o *cpdOptions) run() error {

	// Get the cluster info
	ocmClient := utils.CreateConnection()

	// TODO: This currently only supports the internal OCM ID
	resp, err := ocmClient.ClustersMgmt().V1().Clusters().Cluster(o.clusterID).Get().Send()
	if err != nil {
		return err
	}

	clusterInfo := resp.Body()

	fmt.Println("Checking if cluster has become ready")
	if clusterInfo.Status().State() == "ready" {
		fmt.Printf("This cluster is in a ready state and already provisioned")
		return nil
	}

	fmt.Println("Checking if cluster DNS is ready")
	// Check if DNS is ready, exit out if not
	if !clusterInfo.Status().DNSReady() {
		fmt.Println("DNS not ready. Investigate reasons using the dnszones CR in the cluster namespace:")
		fmt.Printf("oc get dnszones -n uhc-production-%s -o yaml --as backplane-cluster-admin\n", o.clusterID)
		return nil
	}

	fmt.Println("Checking if OCM error code is already known")
	// Check if the OCM Error code is a known error
	if clusterInfo.Status().ProvisionErrorCode() != unknownProvisionCode {
		fmt.Printf("Error code %s is known, customer already received Service Log\n", clusterInfo.Status().ProvisionErrorCode())
	}

	fmt.Println("Checking if cluster is GCP")
	// If the cluster is GCP, give instructions on how to get console access
	if clusterInfo.CloudProvider().ID() == "gcp" {
		fmt.Printf("This command doesn't support GCP yet. Needs manual investigation\n Get the project ID from this command on hive: oc getprojectclaim -n uhc-production-$CLUSTER_INT_ID\nThen use this URL to access the GCP console: https://console.cloud.google.com/home/dashboard?project=${GCP_PROJECT_ID}\n")
		return nil
	}

	fmt.Println("Generating AWS credentials for cluster")
	// Get AWS credentials for the cluster
	awsClient, err := osdCloud.GenerateAWSClientForCluster(o.awsProfile, o.clusterID)
	if err != nil {
		fmt.Println("PLEASE CONFIRM YOUR CREDENTIALS ARE CORRECT. If you're absolutely sure they are, send this Service Log https://github.com/openshift/managed-notifications/blob/master/osd/aws/ROSA_AWS_invalid_permissions.json")
		fmt.Println(err)
		return err
	}

	// If the cluster is BYOVPC, check the route tables
	// This check is copied from ocm-cli
	if clusterInfo.AWS().SubnetIDs() != nil && len(clusterInfo.AWS().SubnetIDs()) > 0 {
		fmt.Println("Checking BYOVPC to ensure subnets have valid routing")
		for _, subnet := range clusterInfo.AWS().SubnetIDs() {
			isValid, err := isSubnetRouteValid(awsClient, subnet)
			if err != nil {
				return err
			}
			if !isValid {
				err := fmt.Errorf("subnet %v does not have valid routing. ", subnet)
				fmt.Printf("%v\n Run the following to send a SerivceLog:\n osdctl servicelog post ${CLUSTER_EXT_ID} -t https://raw.githubusercontent.com/openshift/managed-notifications/master/osd/aws/InstallFailed_NoRouteToInternet.json", err)
				return err
			}
		}
		fmt.Println("Next step: run the verifier egress test: osd-network-verifier egress --region ${CLUSTER_REGION} --subnet-id ${SUBNET_ID} --security-group ${SECURITY_GROUP_ID}")
		return nil
	}

	fmt.Println("Next step: check the AWS reources manually, run ocm backplane cloud console")

	return nil
}

func isSubnetRouteValid(awsClient aws.Client, subnetID string) (bool, error) {

	var routeTable string

	// Try and find a RouteTable with an association for the given subnet
	describeRouteTablesOutput, err := awsClient.DescribeRouteTables(&ec2.DescribeRouteTablesInput{
		Filters: []*ec2.Filter{
			{
				Name:   awsSdk.String("association.subnet-id"),
				Values: []*string{awsSdk.String(subnetID)},
			},
		},
	})
	if err != nil {
		return false, err
	}

	// If there are no associated RouteTables, then the subnet uses the default RoutTable for the VPC
	if len(describeRouteTablesOutput.RouteTables) == 0 {
		// Get the VPC ID for the subnet
		describeSubnetOutput, err := awsClient.DescribeSubnets(&ec2.DescribeSubnetsInput{
			SubnetIds: []*string{&subnetID},
		})
		if err != nil {
			return false, err
		}
		if len(describeSubnetOutput.Subnets) == 0 {
			return false, fmt.Errorf("no subnets returned for subnet id %v", subnetID)
		}

		vpcID := *describeSubnetOutput.Subnets[0].VpcId

		// Set the routetable to the default for the VPC
		routeTable, err = findDefaultRouteTableForVPC(awsClient, vpcID)
		if err != nil {
			return false, err
		}
	} else {
		// Set the routetable to the one asscoiated with the subnet
		routeTable = *describeRouteTablesOutput.RouteTables[0].RouteTableId
	}

	// Check that the RouteTable for the subnet has a default route to 0.0.0.0/0
	describeRouteTablesOutput, err = awsClient.DescribeRouteTables(&ec2.DescribeRouteTablesInput{
		RouteTableIds: []*string{awsSdk.String(routeTable)},
	})
	if err != nil {
		return false, err
	}

	if len(describeRouteTablesOutput.RouteTables) == 0 {
		return false, fmt.Errorf("no routetables found for routetable id %v", routeTable)
	}

	for _, route := range describeRouteTablesOutput.RouteTables[0].Routes {
		// Some routes don't use CIDR blocks as targets, so this needs to be checked
		if route.DestinationCidrBlock != nil {
			if *route.DestinationCidrBlock == "0.0.0.0/0" {
				return true, nil
			}
		}

	}

	return false, fmt.Errorf("no default route exists to the internet in subnet %v", subnetID)
}

func findDefaultRouteTableForVPC(awsClient aws.Client, vpcID string) (string, error) {

	// Get all subnets in the VPC
	describeRouteTablesOutput, err := awsClient.DescribeRouteTables(&ec2.DescribeRouteTablesInput{
		Filters: []*ec2.Filter{
			{
				Name:   awsSdk.String("vpc-id"),
				Values: []*string{awsSdk.String(vpcID)},
			},
		},
	})
	if err != nil {
		return "", err
	}

	for _, rt := range describeRouteTablesOutput.RouteTables {
		for _, assoc := range rt.Associations {
			if *assoc.Main {
				return *rt.RouteTableId, nil
			}
		}
	}

	return "", fmt.Errorf("no default routetable found for vpc %v", vpcID)
}
