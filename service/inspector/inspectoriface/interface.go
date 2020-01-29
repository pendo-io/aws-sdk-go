// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package inspectoriface provides an interface to enable mocking the Amazon Inspector service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package inspectoriface

import (
	"github.com/pendo-io/aws-sdk-go/aws"
	"github.com/pendo-io/aws-sdk-go/aws/request"
	"github.com/pendo-io/aws-sdk-go/service/inspector"
)

// InspectorAPI provides an interface to enable mocking the
// inspector.Inspector service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Inspector.
//    func myFunc(svc inspectoriface.InspectorAPI) bool {
//        // Make svc.AddAttributesToFindings request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := inspector.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockInspectorClient struct {
//        inspectoriface.InspectorAPI
//    }
//    func (m *mockInspectorClient) AddAttributesToFindings(input *inspector.AddAttributesToFindingsInput) (*inspector.AddAttributesToFindingsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockInspectorClient{}
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
type InspectorAPI interface {
	AddAttributesToFindings(*inspector.AddAttributesToFindingsInput) (*inspector.AddAttributesToFindingsOutput, error)
	AddAttributesToFindingsWithContext(aws.Context, *inspector.AddAttributesToFindingsInput, ...request.Option) (*inspector.AddAttributesToFindingsOutput, error)
	AddAttributesToFindingsRequest(*inspector.AddAttributesToFindingsInput) (*request.Request, *inspector.AddAttributesToFindingsOutput)

	CreateAssessmentTarget(*inspector.CreateAssessmentTargetInput) (*inspector.CreateAssessmentTargetOutput, error)
	CreateAssessmentTargetWithContext(aws.Context, *inspector.CreateAssessmentTargetInput, ...request.Option) (*inspector.CreateAssessmentTargetOutput, error)
	CreateAssessmentTargetRequest(*inspector.CreateAssessmentTargetInput) (*request.Request, *inspector.CreateAssessmentTargetOutput)

	CreateAssessmentTemplate(*inspector.CreateAssessmentTemplateInput) (*inspector.CreateAssessmentTemplateOutput, error)
	CreateAssessmentTemplateWithContext(aws.Context, *inspector.CreateAssessmentTemplateInput, ...request.Option) (*inspector.CreateAssessmentTemplateOutput, error)
	CreateAssessmentTemplateRequest(*inspector.CreateAssessmentTemplateInput) (*request.Request, *inspector.CreateAssessmentTemplateOutput)

	CreateExclusionsPreview(*inspector.CreateExclusionsPreviewInput) (*inspector.CreateExclusionsPreviewOutput, error)
	CreateExclusionsPreviewWithContext(aws.Context, *inspector.CreateExclusionsPreviewInput, ...request.Option) (*inspector.CreateExclusionsPreviewOutput, error)
	CreateExclusionsPreviewRequest(*inspector.CreateExclusionsPreviewInput) (*request.Request, *inspector.CreateExclusionsPreviewOutput)

	CreateResourceGroup(*inspector.CreateResourceGroupInput) (*inspector.CreateResourceGroupOutput, error)
	CreateResourceGroupWithContext(aws.Context, *inspector.CreateResourceGroupInput, ...request.Option) (*inspector.CreateResourceGroupOutput, error)
	CreateResourceGroupRequest(*inspector.CreateResourceGroupInput) (*request.Request, *inspector.CreateResourceGroupOutput)

	DeleteAssessmentRun(*inspector.DeleteAssessmentRunInput) (*inspector.DeleteAssessmentRunOutput, error)
	DeleteAssessmentRunWithContext(aws.Context, *inspector.DeleteAssessmentRunInput, ...request.Option) (*inspector.DeleteAssessmentRunOutput, error)
	DeleteAssessmentRunRequest(*inspector.DeleteAssessmentRunInput) (*request.Request, *inspector.DeleteAssessmentRunOutput)

	DeleteAssessmentTarget(*inspector.DeleteAssessmentTargetInput) (*inspector.DeleteAssessmentTargetOutput, error)
	DeleteAssessmentTargetWithContext(aws.Context, *inspector.DeleteAssessmentTargetInput, ...request.Option) (*inspector.DeleteAssessmentTargetOutput, error)
	DeleteAssessmentTargetRequest(*inspector.DeleteAssessmentTargetInput) (*request.Request, *inspector.DeleteAssessmentTargetOutput)

	DeleteAssessmentTemplate(*inspector.DeleteAssessmentTemplateInput) (*inspector.DeleteAssessmentTemplateOutput, error)
	DeleteAssessmentTemplateWithContext(aws.Context, *inspector.DeleteAssessmentTemplateInput, ...request.Option) (*inspector.DeleteAssessmentTemplateOutput, error)
	DeleteAssessmentTemplateRequest(*inspector.DeleteAssessmentTemplateInput) (*request.Request, *inspector.DeleteAssessmentTemplateOutput)

	DescribeAssessmentRuns(*inspector.DescribeAssessmentRunsInput) (*inspector.DescribeAssessmentRunsOutput, error)
	DescribeAssessmentRunsWithContext(aws.Context, *inspector.DescribeAssessmentRunsInput, ...request.Option) (*inspector.DescribeAssessmentRunsOutput, error)
	DescribeAssessmentRunsRequest(*inspector.DescribeAssessmentRunsInput) (*request.Request, *inspector.DescribeAssessmentRunsOutput)

	DescribeAssessmentTargets(*inspector.DescribeAssessmentTargetsInput) (*inspector.DescribeAssessmentTargetsOutput, error)
	DescribeAssessmentTargetsWithContext(aws.Context, *inspector.DescribeAssessmentTargetsInput, ...request.Option) (*inspector.DescribeAssessmentTargetsOutput, error)
	DescribeAssessmentTargetsRequest(*inspector.DescribeAssessmentTargetsInput) (*request.Request, *inspector.DescribeAssessmentTargetsOutput)

	DescribeAssessmentTemplates(*inspector.DescribeAssessmentTemplatesInput) (*inspector.DescribeAssessmentTemplatesOutput, error)
	DescribeAssessmentTemplatesWithContext(aws.Context, *inspector.DescribeAssessmentTemplatesInput, ...request.Option) (*inspector.DescribeAssessmentTemplatesOutput, error)
	DescribeAssessmentTemplatesRequest(*inspector.DescribeAssessmentTemplatesInput) (*request.Request, *inspector.DescribeAssessmentTemplatesOutput)

	DescribeCrossAccountAccessRole(*inspector.DescribeCrossAccountAccessRoleInput) (*inspector.DescribeCrossAccountAccessRoleOutput, error)
	DescribeCrossAccountAccessRoleWithContext(aws.Context, *inspector.DescribeCrossAccountAccessRoleInput, ...request.Option) (*inspector.DescribeCrossAccountAccessRoleOutput, error)
	DescribeCrossAccountAccessRoleRequest(*inspector.DescribeCrossAccountAccessRoleInput) (*request.Request, *inspector.DescribeCrossAccountAccessRoleOutput)

	DescribeExclusions(*inspector.DescribeExclusionsInput) (*inspector.DescribeExclusionsOutput, error)
	DescribeExclusionsWithContext(aws.Context, *inspector.DescribeExclusionsInput, ...request.Option) (*inspector.DescribeExclusionsOutput, error)
	DescribeExclusionsRequest(*inspector.DescribeExclusionsInput) (*request.Request, *inspector.DescribeExclusionsOutput)

	DescribeFindings(*inspector.DescribeFindingsInput) (*inspector.DescribeFindingsOutput, error)
	DescribeFindingsWithContext(aws.Context, *inspector.DescribeFindingsInput, ...request.Option) (*inspector.DescribeFindingsOutput, error)
	DescribeFindingsRequest(*inspector.DescribeFindingsInput) (*request.Request, *inspector.DescribeFindingsOutput)

	DescribeResourceGroups(*inspector.DescribeResourceGroupsInput) (*inspector.DescribeResourceGroupsOutput, error)
	DescribeResourceGroupsWithContext(aws.Context, *inspector.DescribeResourceGroupsInput, ...request.Option) (*inspector.DescribeResourceGroupsOutput, error)
	DescribeResourceGroupsRequest(*inspector.DescribeResourceGroupsInput) (*request.Request, *inspector.DescribeResourceGroupsOutput)

	DescribeRulesPackages(*inspector.DescribeRulesPackagesInput) (*inspector.DescribeRulesPackagesOutput, error)
	DescribeRulesPackagesWithContext(aws.Context, *inspector.DescribeRulesPackagesInput, ...request.Option) (*inspector.DescribeRulesPackagesOutput, error)
	DescribeRulesPackagesRequest(*inspector.DescribeRulesPackagesInput) (*request.Request, *inspector.DescribeRulesPackagesOutput)

	GetAssessmentReport(*inspector.GetAssessmentReportInput) (*inspector.GetAssessmentReportOutput, error)
	GetAssessmentReportWithContext(aws.Context, *inspector.GetAssessmentReportInput, ...request.Option) (*inspector.GetAssessmentReportOutput, error)
	GetAssessmentReportRequest(*inspector.GetAssessmentReportInput) (*request.Request, *inspector.GetAssessmentReportOutput)

	GetExclusionsPreview(*inspector.GetExclusionsPreviewInput) (*inspector.GetExclusionsPreviewOutput, error)
	GetExclusionsPreviewWithContext(aws.Context, *inspector.GetExclusionsPreviewInput, ...request.Option) (*inspector.GetExclusionsPreviewOutput, error)
	GetExclusionsPreviewRequest(*inspector.GetExclusionsPreviewInput) (*request.Request, *inspector.GetExclusionsPreviewOutput)

	GetExclusionsPreviewPages(*inspector.GetExclusionsPreviewInput, func(*inspector.GetExclusionsPreviewOutput, bool) bool) error
	GetExclusionsPreviewPagesWithContext(aws.Context, *inspector.GetExclusionsPreviewInput, func(*inspector.GetExclusionsPreviewOutput, bool) bool, ...request.Option) error

	GetTelemetryMetadata(*inspector.GetTelemetryMetadataInput) (*inspector.GetTelemetryMetadataOutput, error)
	GetTelemetryMetadataWithContext(aws.Context, *inspector.GetTelemetryMetadataInput, ...request.Option) (*inspector.GetTelemetryMetadataOutput, error)
	GetTelemetryMetadataRequest(*inspector.GetTelemetryMetadataInput) (*request.Request, *inspector.GetTelemetryMetadataOutput)

	ListAssessmentRunAgents(*inspector.ListAssessmentRunAgentsInput) (*inspector.ListAssessmentRunAgentsOutput, error)
	ListAssessmentRunAgentsWithContext(aws.Context, *inspector.ListAssessmentRunAgentsInput, ...request.Option) (*inspector.ListAssessmentRunAgentsOutput, error)
	ListAssessmentRunAgentsRequest(*inspector.ListAssessmentRunAgentsInput) (*request.Request, *inspector.ListAssessmentRunAgentsOutput)

	ListAssessmentRunAgentsPages(*inspector.ListAssessmentRunAgentsInput, func(*inspector.ListAssessmentRunAgentsOutput, bool) bool) error
	ListAssessmentRunAgentsPagesWithContext(aws.Context, *inspector.ListAssessmentRunAgentsInput, func(*inspector.ListAssessmentRunAgentsOutput, bool) bool, ...request.Option) error

	ListAssessmentRuns(*inspector.ListAssessmentRunsInput) (*inspector.ListAssessmentRunsOutput, error)
	ListAssessmentRunsWithContext(aws.Context, *inspector.ListAssessmentRunsInput, ...request.Option) (*inspector.ListAssessmentRunsOutput, error)
	ListAssessmentRunsRequest(*inspector.ListAssessmentRunsInput) (*request.Request, *inspector.ListAssessmentRunsOutput)

	ListAssessmentRunsPages(*inspector.ListAssessmentRunsInput, func(*inspector.ListAssessmentRunsOutput, bool) bool) error
	ListAssessmentRunsPagesWithContext(aws.Context, *inspector.ListAssessmentRunsInput, func(*inspector.ListAssessmentRunsOutput, bool) bool, ...request.Option) error

	ListAssessmentTargets(*inspector.ListAssessmentTargetsInput) (*inspector.ListAssessmentTargetsOutput, error)
	ListAssessmentTargetsWithContext(aws.Context, *inspector.ListAssessmentTargetsInput, ...request.Option) (*inspector.ListAssessmentTargetsOutput, error)
	ListAssessmentTargetsRequest(*inspector.ListAssessmentTargetsInput) (*request.Request, *inspector.ListAssessmentTargetsOutput)

	ListAssessmentTargetsPages(*inspector.ListAssessmentTargetsInput, func(*inspector.ListAssessmentTargetsOutput, bool) bool) error
	ListAssessmentTargetsPagesWithContext(aws.Context, *inspector.ListAssessmentTargetsInput, func(*inspector.ListAssessmentTargetsOutput, bool) bool, ...request.Option) error

	ListAssessmentTemplates(*inspector.ListAssessmentTemplatesInput) (*inspector.ListAssessmentTemplatesOutput, error)
	ListAssessmentTemplatesWithContext(aws.Context, *inspector.ListAssessmentTemplatesInput, ...request.Option) (*inspector.ListAssessmentTemplatesOutput, error)
	ListAssessmentTemplatesRequest(*inspector.ListAssessmentTemplatesInput) (*request.Request, *inspector.ListAssessmentTemplatesOutput)

	ListAssessmentTemplatesPages(*inspector.ListAssessmentTemplatesInput, func(*inspector.ListAssessmentTemplatesOutput, bool) bool) error
	ListAssessmentTemplatesPagesWithContext(aws.Context, *inspector.ListAssessmentTemplatesInput, func(*inspector.ListAssessmentTemplatesOutput, bool) bool, ...request.Option) error

	ListEventSubscriptions(*inspector.ListEventSubscriptionsInput) (*inspector.ListEventSubscriptionsOutput, error)
	ListEventSubscriptionsWithContext(aws.Context, *inspector.ListEventSubscriptionsInput, ...request.Option) (*inspector.ListEventSubscriptionsOutput, error)
	ListEventSubscriptionsRequest(*inspector.ListEventSubscriptionsInput) (*request.Request, *inspector.ListEventSubscriptionsOutput)

	ListEventSubscriptionsPages(*inspector.ListEventSubscriptionsInput, func(*inspector.ListEventSubscriptionsOutput, bool) bool) error
	ListEventSubscriptionsPagesWithContext(aws.Context, *inspector.ListEventSubscriptionsInput, func(*inspector.ListEventSubscriptionsOutput, bool) bool, ...request.Option) error

	ListExclusions(*inspector.ListExclusionsInput) (*inspector.ListExclusionsOutput, error)
	ListExclusionsWithContext(aws.Context, *inspector.ListExclusionsInput, ...request.Option) (*inspector.ListExclusionsOutput, error)
	ListExclusionsRequest(*inspector.ListExclusionsInput) (*request.Request, *inspector.ListExclusionsOutput)

	ListExclusionsPages(*inspector.ListExclusionsInput, func(*inspector.ListExclusionsOutput, bool) bool) error
	ListExclusionsPagesWithContext(aws.Context, *inspector.ListExclusionsInput, func(*inspector.ListExclusionsOutput, bool) bool, ...request.Option) error

	ListFindings(*inspector.ListFindingsInput) (*inspector.ListFindingsOutput, error)
	ListFindingsWithContext(aws.Context, *inspector.ListFindingsInput, ...request.Option) (*inspector.ListFindingsOutput, error)
	ListFindingsRequest(*inspector.ListFindingsInput) (*request.Request, *inspector.ListFindingsOutput)

	ListFindingsPages(*inspector.ListFindingsInput, func(*inspector.ListFindingsOutput, bool) bool) error
	ListFindingsPagesWithContext(aws.Context, *inspector.ListFindingsInput, func(*inspector.ListFindingsOutput, bool) bool, ...request.Option) error

	ListRulesPackages(*inspector.ListRulesPackagesInput) (*inspector.ListRulesPackagesOutput, error)
	ListRulesPackagesWithContext(aws.Context, *inspector.ListRulesPackagesInput, ...request.Option) (*inspector.ListRulesPackagesOutput, error)
	ListRulesPackagesRequest(*inspector.ListRulesPackagesInput) (*request.Request, *inspector.ListRulesPackagesOutput)

	ListRulesPackagesPages(*inspector.ListRulesPackagesInput, func(*inspector.ListRulesPackagesOutput, bool) bool) error
	ListRulesPackagesPagesWithContext(aws.Context, *inspector.ListRulesPackagesInput, func(*inspector.ListRulesPackagesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*inspector.ListTagsForResourceInput) (*inspector.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *inspector.ListTagsForResourceInput, ...request.Option) (*inspector.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*inspector.ListTagsForResourceInput) (*request.Request, *inspector.ListTagsForResourceOutput)

	PreviewAgents(*inspector.PreviewAgentsInput) (*inspector.PreviewAgentsOutput, error)
	PreviewAgentsWithContext(aws.Context, *inspector.PreviewAgentsInput, ...request.Option) (*inspector.PreviewAgentsOutput, error)
	PreviewAgentsRequest(*inspector.PreviewAgentsInput) (*request.Request, *inspector.PreviewAgentsOutput)

	PreviewAgentsPages(*inspector.PreviewAgentsInput, func(*inspector.PreviewAgentsOutput, bool) bool) error
	PreviewAgentsPagesWithContext(aws.Context, *inspector.PreviewAgentsInput, func(*inspector.PreviewAgentsOutput, bool) bool, ...request.Option) error

	RegisterCrossAccountAccessRole(*inspector.RegisterCrossAccountAccessRoleInput) (*inspector.RegisterCrossAccountAccessRoleOutput, error)
	RegisterCrossAccountAccessRoleWithContext(aws.Context, *inspector.RegisterCrossAccountAccessRoleInput, ...request.Option) (*inspector.RegisterCrossAccountAccessRoleOutput, error)
	RegisterCrossAccountAccessRoleRequest(*inspector.RegisterCrossAccountAccessRoleInput) (*request.Request, *inspector.RegisterCrossAccountAccessRoleOutput)

	RemoveAttributesFromFindings(*inspector.RemoveAttributesFromFindingsInput) (*inspector.RemoveAttributesFromFindingsOutput, error)
	RemoveAttributesFromFindingsWithContext(aws.Context, *inspector.RemoveAttributesFromFindingsInput, ...request.Option) (*inspector.RemoveAttributesFromFindingsOutput, error)
	RemoveAttributesFromFindingsRequest(*inspector.RemoveAttributesFromFindingsInput) (*request.Request, *inspector.RemoveAttributesFromFindingsOutput)

	SetTagsForResource(*inspector.SetTagsForResourceInput) (*inspector.SetTagsForResourceOutput, error)
	SetTagsForResourceWithContext(aws.Context, *inspector.SetTagsForResourceInput, ...request.Option) (*inspector.SetTagsForResourceOutput, error)
	SetTagsForResourceRequest(*inspector.SetTagsForResourceInput) (*request.Request, *inspector.SetTagsForResourceOutput)

	StartAssessmentRun(*inspector.StartAssessmentRunInput) (*inspector.StartAssessmentRunOutput, error)
	StartAssessmentRunWithContext(aws.Context, *inspector.StartAssessmentRunInput, ...request.Option) (*inspector.StartAssessmentRunOutput, error)
	StartAssessmentRunRequest(*inspector.StartAssessmentRunInput) (*request.Request, *inspector.StartAssessmentRunOutput)

	StopAssessmentRun(*inspector.StopAssessmentRunInput) (*inspector.StopAssessmentRunOutput, error)
	StopAssessmentRunWithContext(aws.Context, *inspector.StopAssessmentRunInput, ...request.Option) (*inspector.StopAssessmentRunOutput, error)
	StopAssessmentRunRequest(*inspector.StopAssessmentRunInput) (*request.Request, *inspector.StopAssessmentRunOutput)

	SubscribeToEvent(*inspector.SubscribeToEventInput) (*inspector.SubscribeToEventOutput, error)
	SubscribeToEventWithContext(aws.Context, *inspector.SubscribeToEventInput, ...request.Option) (*inspector.SubscribeToEventOutput, error)
	SubscribeToEventRequest(*inspector.SubscribeToEventInput) (*request.Request, *inspector.SubscribeToEventOutput)

	UnsubscribeFromEvent(*inspector.UnsubscribeFromEventInput) (*inspector.UnsubscribeFromEventOutput, error)
	UnsubscribeFromEventWithContext(aws.Context, *inspector.UnsubscribeFromEventInput, ...request.Option) (*inspector.UnsubscribeFromEventOutput, error)
	UnsubscribeFromEventRequest(*inspector.UnsubscribeFromEventInput) (*request.Request, *inspector.UnsubscribeFromEventOutput)

	UpdateAssessmentTarget(*inspector.UpdateAssessmentTargetInput) (*inspector.UpdateAssessmentTargetOutput, error)
	UpdateAssessmentTargetWithContext(aws.Context, *inspector.UpdateAssessmentTargetInput, ...request.Option) (*inspector.UpdateAssessmentTargetOutput, error)
	UpdateAssessmentTargetRequest(*inspector.UpdateAssessmentTargetInput) (*request.Request, *inspector.UpdateAssessmentTargetOutput)
}

var _ InspectorAPI = (*inspector.Inspector)(nil)
