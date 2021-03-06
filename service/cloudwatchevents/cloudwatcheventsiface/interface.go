// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudwatcheventsiface provides an interface to enable mocking the Amazon CloudWatch Events service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudwatcheventsiface

import (
	"github.com/pendo-io/aws-sdk-go/aws"
	"github.com/pendo-io/aws-sdk-go/aws/request"
	"github.com/pendo-io/aws-sdk-go/service/cloudwatchevents"
)

// CloudWatchEventsAPI provides an interface to enable mocking the
// cloudwatchevents.CloudWatchEvents service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon CloudWatch Events.
//    func myFunc(svc cloudwatcheventsiface.CloudWatchEventsAPI) bool {
//        // Make svc.ActivateEventSource request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cloudwatchevents.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCloudWatchEventsClient struct {
//        cloudwatcheventsiface.CloudWatchEventsAPI
//    }
//    func (m *mockCloudWatchEventsClient) ActivateEventSource(input *cloudwatchevents.ActivateEventSourceInput) (*cloudwatchevents.ActivateEventSourceOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCloudWatchEventsClient{}
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
type CloudWatchEventsAPI interface {
	ActivateEventSource(*cloudwatchevents.ActivateEventSourceInput) (*cloudwatchevents.ActivateEventSourceOutput, error)
	ActivateEventSourceWithContext(aws.Context, *cloudwatchevents.ActivateEventSourceInput, ...request.Option) (*cloudwatchevents.ActivateEventSourceOutput, error)
	ActivateEventSourceRequest(*cloudwatchevents.ActivateEventSourceInput) (*request.Request, *cloudwatchevents.ActivateEventSourceOutput)

	CreateEventBus(*cloudwatchevents.CreateEventBusInput) (*cloudwatchevents.CreateEventBusOutput, error)
	CreateEventBusWithContext(aws.Context, *cloudwatchevents.CreateEventBusInput, ...request.Option) (*cloudwatchevents.CreateEventBusOutput, error)
	CreateEventBusRequest(*cloudwatchevents.CreateEventBusInput) (*request.Request, *cloudwatchevents.CreateEventBusOutput)

	CreatePartnerEventSource(*cloudwatchevents.CreatePartnerEventSourceInput) (*cloudwatchevents.CreatePartnerEventSourceOutput, error)
	CreatePartnerEventSourceWithContext(aws.Context, *cloudwatchevents.CreatePartnerEventSourceInput, ...request.Option) (*cloudwatchevents.CreatePartnerEventSourceOutput, error)
	CreatePartnerEventSourceRequest(*cloudwatchevents.CreatePartnerEventSourceInput) (*request.Request, *cloudwatchevents.CreatePartnerEventSourceOutput)

	DeactivateEventSource(*cloudwatchevents.DeactivateEventSourceInput) (*cloudwatchevents.DeactivateEventSourceOutput, error)
	DeactivateEventSourceWithContext(aws.Context, *cloudwatchevents.DeactivateEventSourceInput, ...request.Option) (*cloudwatchevents.DeactivateEventSourceOutput, error)
	DeactivateEventSourceRequest(*cloudwatchevents.DeactivateEventSourceInput) (*request.Request, *cloudwatchevents.DeactivateEventSourceOutput)

	DeleteEventBus(*cloudwatchevents.DeleteEventBusInput) (*cloudwatchevents.DeleteEventBusOutput, error)
	DeleteEventBusWithContext(aws.Context, *cloudwatchevents.DeleteEventBusInput, ...request.Option) (*cloudwatchevents.DeleteEventBusOutput, error)
	DeleteEventBusRequest(*cloudwatchevents.DeleteEventBusInput) (*request.Request, *cloudwatchevents.DeleteEventBusOutput)

	DeletePartnerEventSource(*cloudwatchevents.DeletePartnerEventSourceInput) (*cloudwatchevents.DeletePartnerEventSourceOutput, error)
	DeletePartnerEventSourceWithContext(aws.Context, *cloudwatchevents.DeletePartnerEventSourceInput, ...request.Option) (*cloudwatchevents.DeletePartnerEventSourceOutput, error)
	DeletePartnerEventSourceRequest(*cloudwatchevents.DeletePartnerEventSourceInput) (*request.Request, *cloudwatchevents.DeletePartnerEventSourceOutput)

	DeleteRule(*cloudwatchevents.DeleteRuleInput) (*cloudwatchevents.DeleteRuleOutput, error)
	DeleteRuleWithContext(aws.Context, *cloudwatchevents.DeleteRuleInput, ...request.Option) (*cloudwatchevents.DeleteRuleOutput, error)
	DeleteRuleRequest(*cloudwatchevents.DeleteRuleInput) (*request.Request, *cloudwatchevents.DeleteRuleOutput)

	DescribeEventBus(*cloudwatchevents.DescribeEventBusInput) (*cloudwatchevents.DescribeEventBusOutput, error)
	DescribeEventBusWithContext(aws.Context, *cloudwatchevents.DescribeEventBusInput, ...request.Option) (*cloudwatchevents.DescribeEventBusOutput, error)
	DescribeEventBusRequest(*cloudwatchevents.DescribeEventBusInput) (*request.Request, *cloudwatchevents.DescribeEventBusOutput)

	DescribeEventSource(*cloudwatchevents.DescribeEventSourceInput) (*cloudwatchevents.DescribeEventSourceOutput, error)
	DescribeEventSourceWithContext(aws.Context, *cloudwatchevents.DescribeEventSourceInput, ...request.Option) (*cloudwatchevents.DescribeEventSourceOutput, error)
	DescribeEventSourceRequest(*cloudwatchevents.DescribeEventSourceInput) (*request.Request, *cloudwatchevents.DescribeEventSourceOutput)

	DescribePartnerEventSource(*cloudwatchevents.DescribePartnerEventSourceInput) (*cloudwatchevents.DescribePartnerEventSourceOutput, error)
	DescribePartnerEventSourceWithContext(aws.Context, *cloudwatchevents.DescribePartnerEventSourceInput, ...request.Option) (*cloudwatchevents.DescribePartnerEventSourceOutput, error)
	DescribePartnerEventSourceRequest(*cloudwatchevents.DescribePartnerEventSourceInput) (*request.Request, *cloudwatchevents.DescribePartnerEventSourceOutput)

	DescribeRule(*cloudwatchevents.DescribeRuleInput) (*cloudwatchevents.DescribeRuleOutput, error)
	DescribeRuleWithContext(aws.Context, *cloudwatchevents.DescribeRuleInput, ...request.Option) (*cloudwatchevents.DescribeRuleOutput, error)
	DescribeRuleRequest(*cloudwatchevents.DescribeRuleInput) (*request.Request, *cloudwatchevents.DescribeRuleOutput)

	DisableRule(*cloudwatchevents.DisableRuleInput) (*cloudwatchevents.DisableRuleOutput, error)
	DisableRuleWithContext(aws.Context, *cloudwatchevents.DisableRuleInput, ...request.Option) (*cloudwatchevents.DisableRuleOutput, error)
	DisableRuleRequest(*cloudwatchevents.DisableRuleInput) (*request.Request, *cloudwatchevents.DisableRuleOutput)

	EnableRule(*cloudwatchevents.EnableRuleInput) (*cloudwatchevents.EnableRuleOutput, error)
	EnableRuleWithContext(aws.Context, *cloudwatchevents.EnableRuleInput, ...request.Option) (*cloudwatchevents.EnableRuleOutput, error)
	EnableRuleRequest(*cloudwatchevents.EnableRuleInput) (*request.Request, *cloudwatchevents.EnableRuleOutput)

	ListEventBuses(*cloudwatchevents.ListEventBusesInput) (*cloudwatchevents.ListEventBusesOutput, error)
	ListEventBusesWithContext(aws.Context, *cloudwatchevents.ListEventBusesInput, ...request.Option) (*cloudwatchevents.ListEventBusesOutput, error)
	ListEventBusesRequest(*cloudwatchevents.ListEventBusesInput) (*request.Request, *cloudwatchevents.ListEventBusesOutput)

	ListEventSources(*cloudwatchevents.ListEventSourcesInput) (*cloudwatchevents.ListEventSourcesOutput, error)
	ListEventSourcesWithContext(aws.Context, *cloudwatchevents.ListEventSourcesInput, ...request.Option) (*cloudwatchevents.ListEventSourcesOutput, error)
	ListEventSourcesRequest(*cloudwatchevents.ListEventSourcesInput) (*request.Request, *cloudwatchevents.ListEventSourcesOutput)

	ListPartnerEventSourceAccounts(*cloudwatchevents.ListPartnerEventSourceAccountsInput) (*cloudwatchevents.ListPartnerEventSourceAccountsOutput, error)
	ListPartnerEventSourceAccountsWithContext(aws.Context, *cloudwatchevents.ListPartnerEventSourceAccountsInput, ...request.Option) (*cloudwatchevents.ListPartnerEventSourceAccountsOutput, error)
	ListPartnerEventSourceAccountsRequest(*cloudwatchevents.ListPartnerEventSourceAccountsInput) (*request.Request, *cloudwatchevents.ListPartnerEventSourceAccountsOutput)

	ListPartnerEventSources(*cloudwatchevents.ListPartnerEventSourcesInput) (*cloudwatchevents.ListPartnerEventSourcesOutput, error)
	ListPartnerEventSourcesWithContext(aws.Context, *cloudwatchevents.ListPartnerEventSourcesInput, ...request.Option) (*cloudwatchevents.ListPartnerEventSourcesOutput, error)
	ListPartnerEventSourcesRequest(*cloudwatchevents.ListPartnerEventSourcesInput) (*request.Request, *cloudwatchevents.ListPartnerEventSourcesOutput)

	ListRuleNamesByTarget(*cloudwatchevents.ListRuleNamesByTargetInput) (*cloudwatchevents.ListRuleNamesByTargetOutput, error)
	ListRuleNamesByTargetWithContext(aws.Context, *cloudwatchevents.ListRuleNamesByTargetInput, ...request.Option) (*cloudwatchevents.ListRuleNamesByTargetOutput, error)
	ListRuleNamesByTargetRequest(*cloudwatchevents.ListRuleNamesByTargetInput) (*request.Request, *cloudwatchevents.ListRuleNamesByTargetOutput)

	ListRules(*cloudwatchevents.ListRulesInput) (*cloudwatchevents.ListRulesOutput, error)
	ListRulesWithContext(aws.Context, *cloudwatchevents.ListRulesInput, ...request.Option) (*cloudwatchevents.ListRulesOutput, error)
	ListRulesRequest(*cloudwatchevents.ListRulesInput) (*request.Request, *cloudwatchevents.ListRulesOutput)

	ListTagsForResource(*cloudwatchevents.ListTagsForResourceInput) (*cloudwatchevents.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *cloudwatchevents.ListTagsForResourceInput, ...request.Option) (*cloudwatchevents.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*cloudwatchevents.ListTagsForResourceInput) (*request.Request, *cloudwatchevents.ListTagsForResourceOutput)

	ListTargetsByRule(*cloudwatchevents.ListTargetsByRuleInput) (*cloudwatchevents.ListTargetsByRuleOutput, error)
	ListTargetsByRuleWithContext(aws.Context, *cloudwatchevents.ListTargetsByRuleInput, ...request.Option) (*cloudwatchevents.ListTargetsByRuleOutput, error)
	ListTargetsByRuleRequest(*cloudwatchevents.ListTargetsByRuleInput) (*request.Request, *cloudwatchevents.ListTargetsByRuleOutput)

	PutEvents(*cloudwatchevents.PutEventsInput) (*cloudwatchevents.PutEventsOutput, error)
	PutEventsWithContext(aws.Context, *cloudwatchevents.PutEventsInput, ...request.Option) (*cloudwatchevents.PutEventsOutput, error)
	PutEventsRequest(*cloudwatchevents.PutEventsInput) (*request.Request, *cloudwatchevents.PutEventsOutput)

	PutPartnerEvents(*cloudwatchevents.PutPartnerEventsInput) (*cloudwatchevents.PutPartnerEventsOutput, error)
	PutPartnerEventsWithContext(aws.Context, *cloudwatchevents.PutPartnerEventsInput, ...request.Option) (*cloudwatchevents.PutPartnerEventsOutput, error)
	PutPartnerEventsRequest(*cloudwatchevents.PutPartnerEventsInput) (*request.Request, *cloudwatchevents.PutPartnerEventsOutput)

	PutPermission(*cloudwatchevents.PutPermissionInput) (*cloudwatchevents.PutPermissionOutput, error)
	PutPermissionWithContext(aws.Context, *cloudwatchevents.PutPermissionInput, ...request.Option) (*cloudwatchevents.PutPermissionOutput, error)
	PutPermissionRequest(*cloudwatchevents.PutPermissionInput) (*request.Request, *cloudwatchevents.PutPermissionOutput)

	PutRule(*cloudwatchevents.PutRuleInput) (*cloudwatchevents.PutRuleOutput, error)
	PutRuleWithContext(aws.Context, *cloudwatchevents.PutRuleInput, ...request.Option) (*cloudwatchevents.PutRuleOutput, error)
	PutRuleRequest(*cloudwatchevents.PutRuleInput) (*request.Request, *cloudwatchevents.PutRuleOutput)

	PutTargets(*cloudwatchevents.PutTargetsInput) (*cloudwatchevents.PutTargetsOutput, error)
	PutTargetsWithContext(aws.Context, *cloudwatchevents.PutTargetsInput, ...request.Option) (*cloudwatchevents.PutTargetsOutput, error)
	PutTargetsRequest(*cloudwatchevents.PutTargetsInput) (*request.Request, *cloudwatchevents.PutTargetsOutput)

	RemovePermission(*cloudwatchevents.RemovePermissionInput) (*cloudwatchevents.RemovePermissionOutput, error)
	RemovePermissionWithContext(aws.Context, *cloudwatchevents.RemovePermissionInput, ...request.Option) (*cloudwatchevents.RemovePermissionOutput, error)
	RemovePermissionRequest(*cloudwatchevents.RemovePermissionInput) (*request.Request, *cloudwatchevents.RemovePermissionOutput)

	RemoveTargets(*cloudwatchevents.RemoveTargetsInput) (*cloudwatchevents.RemoveTargetsOutput, error)
	RemoveTargetsWithContext(aws.Context, *cloudwatchevents.RemoveTargetsInput, ...request.Option) (*cloudwatchevents.RemoveTargetsOutput, error)
	RemoveTargetsRequest(*cloudwatchevents.RemoveTargetsInput) (*request.Request, *cloudwatchevents.RemoveTargetsOutput)

	TagResource(*cloudwatchevents.TagResourceInput) (*cloudwatchevents.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *cloudwatchevents.TagResourceInput, ...request.Option) (*cloudwatchevents.TagResourceOutput, error)
	TagResourceRequest(*cloudwatchevents.TagResourceInput) (*request.Request, *cloudwatchevents.TagResourceOutput)

	TestEventPattern(*cloudwatchevents.TestEventPatternInput) (*cloudwatchevents.TestEventPatternOutput, error)
	TestEventPatternWithContext(aws.Context, *cloudwatchevents.TestEventPatternInput, ...request.Option) (*cloudwatchevents.TestEventPatternOutput, error)
	TestEventPatternRequest(*cloudwatchevents.TestEventPatternInput) (*request.Request, *cloudwatchevents.TestEventPatternOutput)

	UntagResource(*cloudwatchevents.UntagResourceInput) (*cloudwatchevents.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *cloudwatchevents.UntagResourceInput, ...request.Option) (*cloudwatchevents.UntagResourceOutput, error)
	UntagResourceRequest(*cloudwatchevents.UntagResourceInput) (*request.Request, *cloudwatchevents.UntagResourceOutput)
}

var _ CloudWatchEventsAPI = (*cloudwatchevents.CloudWatchEvents)(nil)
