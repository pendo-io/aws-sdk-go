// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package fmsiface provides an interface to enable mocking the Firewall Management Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package fmsiface

import (
	"github.com/pendo-io/aws-sdk-go/aws"
	"github.com/pendo-io/aws-sdk-go/aws/request"
	"github.com/pendo-io/aws-sdk-go/service/fms"
)

// FMSAPI provides an interface to enable mocking the
// fms.FMS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Firewall Management Service.
//    func myFunc(svc fmsiface.FMSAPI) bool {
//        // Make svc.AssociateAdminAccount request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := fms.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockFMSClient struct {
//        fmsiface.FMSAPI
//    }
//    func (m *mockFMSClient) AssociateAdminAccount(input *fms.AssociateAdminAccountInput) (*fms.AssociateAdminAccountOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockFMSClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type FMSAPI interface {
	AssociateAdminAccount(*fms.AssociateAdminAccountInput) (*fms.AssociateAdminAccountOutput, error)
	AssociateAdminAccountWithContext(aws.Context, *fms.AssociateAdminAccountInput, ...request.Option) (*fms.AssociateAdminAccountOutput, error)
	AssociateAdminAccountRequest(*fms.AssociateAdminAccountInput) (*request.Request, *fms.AssociateAdminAccountOutput)

	DeleteNotificationChannel(*fms.DeleteNotificationChannelInput) (*fms.DeleteNotificationChannelOutput, error)
	DeleteNotificationChannelWithContext(aws.Context, *fms.DeleteNotificationChannelInput, ...request.Option) (*fms.DeleteNotificationChannelOutput, error)
	DeleteNotificationChannelRequest(*fms.DeleteNotificationChannelInput) (*request.Request, *fms.DeleteNotificationChannelOutput)

	DeletePolicy(*fms.DeletePolicyInput) (*fms.DeletePolicyOutput, error)
	DeletePolicyWithContext(aws.Context, *fms.DeletePolicyInput, ...request.Option) (*fms.DeletePolicyOutput, error)
	DeletePolicyRequest(*fms.DeletePolicyInput) (*request.Request, *fms.DeletePolicyOutput)

	DisassociateAdminAccount(*fms.DisassociateAdminAccountInput) (*fms.DisassociateAdminAccountOutput, error)
	DisassociateAdminAccountWithContext(aws.Context, *fms.DisassociateAdminAccountInput, ...request.Option) (*fms.DisassociateAdminAccountOutput, error)
	DisassociateAdminAccountRequest(*fms.DisassociateAdminAccountInput) (*request.Request, *fms.DisassociateAdminAccountOutput)

	GetAdminAccount(*fms.GetAdminAccountInput) (*fms.GetAdminAccountOutput, error)
	GetAdminAccountWithContext(aws.Context, *fms.GetAdminAccountInput, ...request.Option) (*fms.GetAdminAccountOutput, error)
	GetAdminAccountRequest(*fms.GetAdminAccountInput) (*request.Request, *fms.GetAdminAccountOutput)

	GetComplianceDetail(*fms.GetComplianceDetailInput) (*fms.GetComplianceDetailOutput, error)
	GetComplianceDetailWithContext(aws.Context, *fms.GetComplianceDetailInput, ...request.Option) (*fms.GetComplianceDetailOutput, error)
	GetComplianceDetailRequest(*fms.GetComplianceDetailInput) (*request.Request, *fms.GetComplianceDetailOutput)

	GetNotificationChannel(*fms.GetNotificationChannelInput) (*fms.GetNotificationChannelOutput, error)
	GetNotificationChannelWithContext(aws.Context, *fms.GetNotificationChannelInput, ...request.Option) (*fms.GetNotificationChannelOutput, error)
	GetNotificationChannelRequest(*fms.GetNotificationChannelInput) (*request.Request, *fms.GetNotificationChannelOutput)

	GetPolicy(*fms.GetPolicyInput) (*fms.GetPolicyOutput, error)
	GetPolicyWithContext(aws.Context, *fms.GetPolicyInput, ...request.Option) (*fms.GetPolicyOutput, error)
	GetPolicyRequest(*fms.GetPolicyInput) (*request.Request, *fms.GetPolicyOutput)

	GetProtectionStatus(*fms.GetProtectionStatusInput) (*fms.GetProtectionStatusOutput, error)
	GetProtectionStatusWithContext(aws.Context, *fms.GetProtectionStatusInput, ...request.Option) (*fms.GetProtectionStatusOutput, error)
	GetProtectionStatusRequest(*fms.GetProtectionStatusInput) (*request.Request, *fms.GetProtectionStatusOutput)

	ListComplianceStatus(*fms.ListComplianceStatusInput) (*fms.ListComplianceStatusOutput, error)
	ListComplianceStatusWithContext(aws.Context, *fms.ListComplianceStatusInput, ...request.Option) (*fms.ListComplianceStatusOutput, error)
	ListComplianceStatusRequest(*fms.ListComplianceStatusInput) (*request.Request, *fms.ListComplianceStatusOutput)

	ListComplianceStatusPages(*fms.ListComplianceStatusInput, func(*fms.ListComplianceStatusOutput, bool) bool) error
	ListComplianceStatusPagesWithContext(aws.Context, *fms.ListComplianceStatusInput, func(*fms.ListComplianceStatusOutput, bool) bool, ...request.Option) error

	ListMemberAccounts(*fms.ListMemberAccountsInput) (*fms.ListMemberAccountsOutput, error)
	ListMemberAccountsWithContext(aws.Context, *fms.ListMemberAccountsInput, ...request.Option) (*fms.ListMemberAccountsOutput, error)
	ListMemberAccountsRequest(*fms.ListMemberAccountsInput) (*request.Request, *fms.ListMemberAccountsOutput)

	ListMemberAccountsPages(*fms.ListMemberAccountsInput, func(*fms.ListMemberAccountsOutput, bool) bool) error
	ListMemberAccountsPagesWithContext(aws.Context, *fms.ListMemberAccountsInput, func(*fms.ListMemberAccountsOutput, bool) bool, ...request.Option) error

	ListPolicies(*fms.ListPoliciesInput) (*fms.ListPoliciesOutput, error)
	ListPoliciesWithContext(aws.Context, *fms.ListPoliciesInput, ...request.Option) (*fms.ListPoliciesOutput, error)
	ListPoliciesRequest(*fms.ListPoliciesInput) (*request.Request, *fms.ListPoliciesOutput)

	ListPoliciesPages(*fms.ListPoliciesInput, func(*fms.ListPoliciesOutput, bool) bool) error
	ListPoliciesPagesWithContext(aws.Context, *fms.ListPoliciesInput, func(*fms.ListPoliciesOutput, bool) bool, ...request.Option) error

	PutNotificationChannel(*fms.PutNotificationChannelInput) (*fms.PutNotificationChannelOutput, error)
	PutNotificationChannelWithContext(aws.Context, *fms.PutNotificationChannelInput, ...request.Option) (*fms.PutNotificationChannelOutput, error)
	PutNotificationChannelRequest(*fms.PutNotificationChannelInput) (*request.Request, *fms.PutNotificationChannelOutput)

	PutPolicy(*fms.PutPolicyInput) (*fms.PutPolicyOutput, error)
	PutPolicyWithContext(aws.Context, *fms.PutPolicyInput, ...request.Option) (*fms.PutPolicyOutput, error)
	PutPolicyRequest(*fms.PutPolicyInput) (*request.Request, *fms.PutPolicyOutput)
}

var _ FMSAPI = (*fms.FMS)(nil)
