// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock is a generated GoMock package.
package mock

import (
	iam "github.com/aws/aws-sdk-go/service/iam"
	s3 "github.com/aws/aws-sdk-go/service/s3"
	sts "github.com/aws/aws-sdk-go/service/sts"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AssumeRole mocks base method.
func (m *MockClient) AssumeRole(arg0 *sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssumeRole", arg0)
	ret0, _ := ret[0].(*sts.AssumeRoleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssumeRole indicates an expected call of AssumeRole.
func (mr *MockClientMockRecorder) AssumeRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRole", reflect.TypeOf((*MockClient)(nil).AssumeRole), arg0)
}

// GetCallerIdentity mocks base method.
func (m *MockClient) GetCallerIdentity(arg0 *sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCallerIdentity", arg0)
	ret0, _ := ret[0].(*sts.GetCallerIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCallerIdentity indicates an expected call of GetCallerIdentity.
func (mr *MockClientMockRecorder) GetCallerIdentity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCallerIdentity", reflect.TypeOf((*MockClient)(nil).GetCallerIdentity), arg0)
}

// GetFederationToken mocks base method.
func (m *MockClient) GetFederationToken(arg0 *sts.GetFederationTokenInput) (*sts.GetFederationTokenOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFederationToken", arg0)
	ret0, _ := ret[0].(*sts.GetFederationTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFederationToken indicates an expected call of GetFederationToken.
func (mr *MockClientMockRecorder) GetFederationToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFederationToken", reflect.TypeOf((*MockClient)(nil).GetFederationToken), arg0)
}

// ListBuckets mocks base method.
func (m *MockClient) ListBuckets(arg0 *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBuckets", arg0)
	ret0, _ := ret[0].(*s3.ListBucketsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBuckets indicates an expected call of ListBuckets.
func (mr *MockClientMockRecorder) ListBuckets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBuckets", reflect.TypeOf((*MockClient)(nil).ListBuckets), arg0)
}

// DeleteBucket mocks base method.
func (m *MockClient) DeleteBucket(arg0 *s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBucket", arg0)
	ret0, _ := ret[0].(*s3.DeleteBucketOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBucket indicates an expected call of DeleteBucket.
func (mr *MockClientMockRecorder) DeleteBucket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBucket", reflect.TypeOf((*MockClient)(nil).DeleteBucket), arg0)
}

// ListObjects mocks base method.
func (m *MockClient) ListObjects(arg0 *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListObjects", arg0)
	ret0, _ := ret[0].(*s3.ListObjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListObjects indicates an expected call of ListObjects.
func (mr *MockClientMockRecorder) ListObjects(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjects", reflect.TypeOf((*MockClient)(nil).ListObjects), arg0)
}

// DeleteObjects mocks base method.
func (m *MockClient) DeleteObjects(arg0 *s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObjects", arg0)
	ret0, _ := ret[0].(*s3.DeleteObjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteObjects indicates an expected call of DeleteObjects.
func (mr *MockClientMockRecorder) DeleteObjects(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObjects", reflect.TypeOf((*MockClient)(nil).DeleteObjects), arg0)
}

// CreateAccessKey mocks base method.
func (m *MockClient) CreateAccessKey(arg0 *iam.CreateAccessKeyInput) (*iam.CreateAccessKeyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessKey", arg0)
	ret0, _ := ret[0].(*iam.CreateAccessKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessKey indicates an expected call of CreateAccessKey.
func (mr *MockClientMockRecorder) CreateAccessKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessKey", reflect.TypeOf((*MockClient)(nil).CreateAccessKey), arg0)
}

// DeleteAccessKey mocks base method.
func (m *MockClient) DeleteAccessKey(arg0 *iam.DeleteAccessKeyInput) (*iam.DeleteAccessKeyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccessKey", arg0)
	ret0, _ := ret[0].(*iam.DeleteAccessKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAccessKey indicates an expected call of DeleteAccessKey.
func (mr *MockClientMockRecorder) DeleteAccessKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccessKey", reflect.TypeOf((*MockClient)(nil).DeleteAccessKey), arg0)
}

// ListAccessKeys mocks base method.
func (m *MockClient) ListAccessKeys(arg0 *iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccessKeys", arg0)
	ret0, _ := ret[0].(*iam.ListAccessKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccessKeys indicates an expected call of ListAccessKeys.
func (mr *MockClientMockRecorder) ListAccessKeys(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessKeys", reflect.TypeOf((*MockClient)(nil).ListAccessKeys), arg0)
}

// GetUser mocks base method.
func (m *MockClient) GetUser(arg0 *iam.GetUserInput) (*iam.GetUserOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(*iam.GetUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockClientMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockClient)(nil).GetUser), arg0)
}

// CreateUser mocks base method.
func (m *MockClient) CreateUser(arg0 *iam.CreateUserInput) (*iam.CreateUserOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(*iam.CreateUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockClientMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockClient)(nil).CreateUser), arg0)
}

// ListUsers mocks base method.
func (m *MockClient) ListUsers(arg0 *iam.ListUsersInput) (*iam.ListUsersOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0)
	ret0, _ := ret[0].(*iam.ListUsersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockClientMockRecorder) ListUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockClient)(nil).ListUsers), arg0)
}

// AttachUserPolicy mocks base method.
func (m *MockClient) AttachUserPolicy(arg0 *iam.AttachUserPolicyInput) (*iam.AttachUserPolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachUserPolicy", arg0)
	ret0, _ := ret[0].(*iam.AttachUserPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachUserPolicy indicates an expected call of AttachUserPolicy.
func (mr *MockClientMockRecorder) AttachUserPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachUserPolicy", reflect.TypeOf((*MockClient)(nil).AttachUserPolicy), arg0)
}

// CreatePolicy mocks base method.
func (m *MockClient) CreatePolicy(arg0 *iam.CreatePolicyInput) (*iam.CreatePolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePolicy", arg0)
	ret0, _ := ret[0].(*iam.CreatePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePolicy indicates an expected call of CreatePolicy.
func (mr *MockClientMockRecorder) CreatePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePolicy", reflect.TypeOf((*MockClient)(nil).CreatePolicy), arg0)
}

// DeletePolicy mocks base method.
func (m *MockClient) DeletePolicy(arg0 *iam.DeletePolicyInput) (*iam.DeletePolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePolicy", arg0)
	ret0, _ := ret[0].(*iam.DeletePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePolicy indicates an expected call of DeletePolicy.
func (mr *MockClientMockRecorder) DeletePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePolicy", reflect.TypeOf((*MockClient)(nil).DeletePolicy), arg0)
}

// AttachRolePolicy mocks base method.
func (m *MockClient) AttachRolePolicy(arg0 *iam.AttachRolePolicyInput) (*iam.AttachRolePolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachRolePolicy", arg0)
	ret0, _ := ret[0].(*iam.AttachRolePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachRolePolicy indicates an expected call of AttachRolePolicy.
func (mr *MockClientMockRecorder) AttachRolePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachRolePolicy", reflect.TypeOf((*MockClient)(nil).AttachRolePolicy), arg0)
}

// DetachRolePolicy mocks base method.
func (m *MockClient) DetachRolePolicy(arg0 *iam.DetachRolePolicyInput) (*iam.DetachRolePolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachRolePolicy", arg0)
	ret0, _ := ret[0].(*iam.DetachRolePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetachRolePolicy indicates an expected call of DetachRolePolicy.
func (mr *MockClientMockRecorder) DetachRolePolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachRolePolicy", reflect.TypeOf((*MockClient)(nil).DetachRolePolicy), arg0)
}

// ListAttachedRolePolicies mocks base method.
func (m *MockClient) ListAttachedRolePolicies(arg0 *iam.ListAttachedRolePoliciesInput) (*iam.ListAttachedRolePoliciesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAttachedRolePolicies", arg0)
	ret0, _ := ret[0].(*iam.ListAttachedRolePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAttachedRolePolicies indicates an expected call of ListAttachedRolePolicies.
func (mr *MockClientMockRecorder) ListAttachedRolePolicies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttachedRolePolicies", reflect.TypeOf((*MockClient)(nil).ListAttachedRolePolicies), arg0)
}
