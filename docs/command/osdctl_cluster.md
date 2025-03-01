## osdctl cluster

Provides information for a specified cluster

### Options

```
  -h, --help   help for cluster
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --as string                        Username to impersonate for the operation. User could be a regular user or a service account in a namespace.
      --cluster string                   The name of the kubeconfig cluster to use
      --context string                   The name of the kubeconfig context to use
      --insecure-skip-tls-verify         If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure
      --kubeconfig string                Path to the kubeconfig file to use for CLI requests.
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files
  -o, --output string                    Valid formats are ['', 'json', 'yaml', 'env']
      --request-timeout string           The length of time to wait before giving up on a single server request. Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests. (default "0")
  -s, --server string                    The address and port of the Kubernetes API server
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO

* [osdctl](osdctl.md)	 - OSD CLI
* [osdctl cluster break-glass](osdctl_cluster_break-glass.md)	 - Emergency access to a cluster
* [osdctl cluster context](osdctl_cluster_context.md)	 - Shows the context of a specified cluster
* [osdctl cluster health](osdctl_cluster_health.md)	 - Describes health of cluster nodes and provides other cluster vitals.
* [osdctl cluster logging-check](osdctl_cluster_logging-check.md)	 - Shows the logging support status of a specified cluster
* [osdctl cluster owner](osdctl_cluster_owner.md)	 - List the clusters owned by the user (can be specified to any user, not only yourself)
* [osdctl cluster resize-control-plane-node](osdctl_cluster_resize-control-plane-node.md)	 - Resize a control plane node. Requires previous login to the api server via `ocm login` and being tunneled to the backplane.
* [osdctl cluster support](osdctl_cluster_support.md)	 - Cluster Support
* [osdctl cluster transfer-owner](osdctl_cluster_transfer-owner.md)	 - Transfer cluster ownership to a new user (to be done by Region Lead)

