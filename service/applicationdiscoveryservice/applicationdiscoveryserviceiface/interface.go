// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package applicationdiscoveryserviceiface provides an interface to enable mocking the AWS Application Discovery Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package applicationdiscoveryserviceiface

import (
	"github.com/pendo-io/aws-sdk-go/aws"
	"github.com/pendo-io/aws-sdk-go/aws/request"
	"github.com/pendo-io/aws-sdk-go/service/applicationdiscoveryservice"
)

// ApplicationDiscoveryServiceAPI provides an interface to enable mocking the
// applicationdiscoveryservice.ApplicationDiscoveryService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Application Discovery Service.
//    func myFunc(svc applicationdiscoveryserviceiface.ApplicationDiscoveryServiceAPI) bool {
//        // Make svc.AssociateConfigurationItemsToApplication request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := applicationdiscoveryservice.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockApplicationDiscoveryServiceClient struct {
//        applicationdiscoveryserviceiface.ApplicationDiscoveryServiceAPI
//    }
//    func (m *mockApplicationDiscoveryServiceClient) AssociateConfigurationItemsToApplication(input *applicationdiscoveryservice.AssociateConfigurationItemsToApplicationInput) (*applicationdiscoveryservice.AssociateConfigurationItemsToApplicationOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockApplicationDiscoveryServiceClient{}
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
type ApplicationDiscoveryServiceAPI interface {
	AssociateConfigurationItemsToApplication(*applicationdiscoveryservice.AssociateConfigurationItemsToApplicationInput) (*applicationdiscoveryservice.AssociateConfigurationItemsToApplicationOutput, error)
	AssociateConfigurationItemsToApplicationWithContext(aws.Context, *applicationdiscoveryservice.AssociateConfigurationItemsToApplicationInput, ...request.Option) (*applicationdiscoveryservice.AssociateConfigurationItemsToApplicationOutput, error)
	AssociateConfigurationItemsToApplicationRequest(*applicationdiscoveryservice.AssociateConfigurationItemsToApplicationInput) (*request.Request, *applicationdiscoveryservice.AssociateConfigurationItemsToApplicationOutput)

	BatchDeleteImportData(*applicationdiscoveryservice.BatchDeleteImportDataInput) (*applicationdiscoveryservice.BatchDeleteImportDataOutput, error)
	BatchDeleteImportDataWithContext(aws.Context, *applicationdiscoveryservice.BatchDeleteImportDataInput, ...request.Option) (*applicationdiscoveryservice.BatchDeleteImportDataOutput, error)
	BatchDeleteImportDataRequest(*applicationdiscoveryservice.BatchDeleteImportDataInput) (*request.Request, *applicationdiscoveryservice.BatchDeleteImportDataOutput)

	CreateApplication(*applicationdiscoveryservice.CreateApplicationInput) (*applicationdiscoveryservice.CreateApplicationOutput, error)
	CreateApplicationWithContext(aws.Context, *applicationdiscoveryservice.CreateApplicationInput, ...request.Option) (*applicationdiscoveryservice.CreateApplicationOutput, error)
	CreateApplicationRequest(*applicationdiscoveryservice.CreateApplicationInput) (*request.Request, *applicationdiscoveryservice.CreateApplicationOutput)

	CreateTags(*applicationdiscoveryservice.CreateTagsInput) (*applicationdiscoveryservice.CreateTagsOutput, error)
	CreateTagsWithContext(aws.Context, *applicationdiscoveryservice.CreateTagsInput, ...request.Option) (*applicationdiscoveryservice.CreateTagsOutput, error)
	CreateTagsRequest(*applicationdiscoveryservice.CreateTagsInput) (*request.Request, *applicationdiscoveryservice.CreateTagsOutput)

	DeleteApplications(*applicationdiscoveryservice.DeleteApplicationsInput) (*applicationdiscoveryservice.DeleteApplicationsOutput, error)
	DeleteApplicationsWithContext(aws.Context, *applicationdiscoveryservice.DeleteApplicationsInput, ...request.Option) (*applicationdiscoveryservice.DeleteApplicationsOutput, error)
	DeleteApplicationsRequest(*applicationdiscoveryservice.DeleteApplicationsInput) (*request.Request, *applicationdiscoveryservice.DeleteApplicationsOutput)

	DeleteTags(*applicationdiscoveryservice.DeleteTagsInput) (*applicationdiscoveryservice.DeleteTagsOutput, error)
	DeleteTagsWithContext(aws.Context, *applicationdiscoveryservice.DeleteTagsInput, ...request.Option) (*applicationdiscoveryservice.DeleteTagsOutput, error)
	DeleteTagsRequest(*applicationdiscoveryservice.DeleteTagsInput) (*request.Request, *applicationdiscoveryservice.DeleteTagsOutput)

	DescribeAgents(*applicationdiscoveryservice.DescribeAgentsInput) (*applicationdiscoveryservice.DescribeAgentsOutput, error)
	DescribeAgentsWithContext(aws.Context, *applicationdiscoveryservice.DescribeAgentsInput, ...request.Option) (*applicationdiscoveryservice.DescribeAgentsOutput, error)
	DescribeAgentsRequest(*applicationdiscoveryservice.DescribeAgentsInput) (*request.Request, *applicationdiscoveryservice.DescribeAgentsOutput)

	DescribeConfigurations(*applicationdiscoveryservice.DescribeConfigurationsInput) (*applicationdiscoveryservice.DescribeConfigurationsOutput, error)
	DescribeConfigurationsWithContext(aws.Context, *applicationdiscoveryservice.DescribeConfigurationsInput, ...request.Option) (*applicationdiscoveryservice.DescribeConfigurationsOutput, error)
	DescribeConfigurationsRequest(*applicationdiscoveryservice.DescribeConfigurationsInput) (*request.Request, *applicationdiscoveryservice.DescribeConfigurationsOutput)

	DescribeContinuousExports(*applicationdiscoveryservice.DescribeContinuousExportsInput) (*applicationdiscoveryservice.DescribeContinuousExportsOutput, error)
	DescribeContinuousExportsWithContext(aws.Context, *applicationdiscoveryservice.DescribeContinuousExportsInput, ...request.Option) (*applicationdiscoveryservice.DescribeContinuousExportsOutput, error)
	DescribeContinuousExportsRequest(*applicationdiscoveryservice.DescribeContinuousExportsInput) (*request.Request, *applicationdiscoveryservice.DescribeContinuousExportsOutput)

	DescribeContinuousExportsPages(*applicationdiscoveryservice.DescribeContinuousExportsInput, func(*applicationdiscoveryservice.DescribeContinuousExportsOutput, bool) bool) error
	DescribeContinuousExportsPagesWithContext(aws.Context, *applicationdiscoveryservice.DescribeContinuousExportsInput, func(*applicationdiscoveryservice.DescribeContinuousExportsOutput, bool) bool, ...request.Option) error

	DescribeExportConfigurations(*applicationdiscoveryservice.DescribeExportConfigurationsInput) (*applicationdiscoveryservice.DescribeExportConfigurationsOutput, error)
	DescribeExportConfigurationsWithContext(aws.Context, *applicationdiscoveryservice.DescribeExportConfigurationsInput, ...request.Option) (*applicationdiscoveryservice.DescribeExportConfigurationsOutput, error)
	DescribeExportConfigurationsRequest(*applicationdiscoveryservice.DescribeExportConfigurationsInput) (*request.Request, *applicationdiscoveryservice.DescribeExportConfigurationsOutput)

	DescribeExportTasks(*applicationdiscoveryservice.DescribeExportTasksInput) (*applicationdiscoveryservice.DescribeExportTasksOutput, error)
	DescribeExportTasksWithContext(aws.Context, *applicationdiscoveryservice.DescribeExportTasksInput, ...request.Option) (*applicationdiscoveryservice.DescribeExportTasksOutput, error)
	DescribeExportTasksRequest(*applicationdiscoveryservice.DescribeExportTasksInput) (*request.Request, *applicationdiscoveryservice.DescribeExportTasksOutput)

	DescribeImportTasks(*applicationdiscoveryservice.DescribeImportTasksInput) (*applicationdiscoveryservice.DescribeImportTasksOutput, error)
	DescribeImportTasksWithContext(aws.Context, *applicationdiscoveryservice.DescribeImportTasksInput, ...request.Option) (*applicationdiscoveryservice.DescribeImportTasksOutput, error)
	DescribeImportTasksRequest(*applicationdiscoveryservice.DescribeImportTasksInput) (*request.Request, *applicationdiscoveryservice.DescribeImportTasksOutput)

	DescribeImportTasksPages(*applicationdiscoveryservice.DescribeImportTasksInput, func(*applicationdiscoveryservice.DescribeImportTasksOutput, bool) bool) error
	DescribeImportTasksPagesWithContext(aws.Context, *applicationdiscoveryservice.DescribeImportTasksInput, func(*applicationdiscoveryservice.DescribeImportTasksOutput, bool) bool, ...request.Option) error

	DescribeTags(*applicationdiscoveryservice.DescribeTagsInput) (*applicationdiscoveryservice.DescribeTagsOutput, error)
	DescribeTagsWithContext(aws.Context, *applicationdiscoveryservice.DescribeTagsInput, ...request.Option) (*applicationdiscoveryservice.DescribeTagsOutput, error)
	DescribeTagsRequest(*applicationdiscoveryservice.DescribeTagsInput) (*request.Request, *applicationdiscoveryservice.DescribeTagsOutput)

	DisassociateConfigurationItemsFromApplication(*applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationInput) (*applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationOutput, error)
	DisassociateConfigurationItemsFromApplicationWithContext(aws.Context, *applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationInput, ...request.Option) (*applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationOutput, error)
	DisassociateConfigurationItemsFromApplicationRequest(*applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationInput) (*request.Request, *applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationOutput)

	ExportConfigurations(*applicationdiscoveryservice.ExportConfigurationsInput) (*applicationdiscoveryservice.ExportConfigurationsOutput, error)
	ExportConfigurationsWithContext(aws.Context, *applicationdiscoveryservice.ExportConfigurationsInput, ...request.Option) (*applicationdiscoveryservice.ExportConfigurationsOutput, error)
	ExportConfigurationsRequest(*applicationdiscoveryservice.ExportConfigurationsInput) (*request.Request, *applicationdiscoveryservice.ExportConfigurationsOutput)

	GetDiscoverySummary(*applicationdiscoveryservice.GetDiscoverySummaryInput) (*applicationdiscoveryservice.GetDiscoverySummaryOutput, error)
	GetDiscoverySummaryWithContext(aws.Context, *applicationdiscoveryservice.GetDiscoverySummaryInput, ...request.Option) (*applicationdiscoveryservice.GetDiscoverySummaryOutput, error)
	GetDiscoverySummaryRequest(*applicationdiscoveryservice.GetDiscoverySummaryInput) (*request.Request, *applicationdiscoveryservice.GetDiscoverySummaryOutput)

	ListConfigurations(*applicationdiscoveryservice.ListConfigurationsInput) (*applicationdiscoveryservice.ListConfigurationsOutput, error)
	ListConfigurationsWithContext(aws.Context, *applicationdiscoveryservice.ListConfigurationsInput, ...request.Option) (*applicationdiscoveryservice.ListConfigurationsOutput, error)
	ListConfigurationsRequest(*applicationdiscoveryservice.ListConfigurationsInput) (*request.Request, *applicationdiscoveryservice.ListConfigurationsOutput)

	ListServerNeighbors(*applicationdiscoveryservice.ListServerNeighborsInput) (*applicationdiscoveryservice.ListServerNeighborsOutput, error)
	ListServerNeighborsWithContext(aws.Context, *applicationdiscoveryservice.ListServerNeighborsInput, ...request.Option) (*applicationdiscoveryservice.ListServerNeighborsOutput, error)
	ListServerNeighborsRequest(*applicationdiscoveryservice.ListServerNeighborsInput) (*request.Request, *applicationdiscoveryservice.ListServerNeighborsOutput)

	StartContinuousExport(*applicationdiscoveryservice.StartContinuousExportInput) (*applicationdiscoveryservice.StartContinuousExportOutput, error)
	StartContinuousExportWithContext(aws.Context, *applicationdiscoveryservice.StartContinuousExportInput, ...request.Option) (*applicationdiscoveryservice.StartContinuousExportOutput, error)
	StartContinuousExportRequest(*applicationdiscoveryservice.StartContinuousExportInput) (*request.Request, *applicationdiscoveryservice.StartContinuousExportOutput)

	StartDataCollectionByAgentIds(*applicationdiscoveryservice.StartDataCollectionByAgentIdsInput) (*applicationdiscoveryservice.StartDataCollectionByAgentIdsOutput, error)
	StartDataCollectionByAgentIdsWithContext(aws.Context, *applicationdiscoveryservice.StartDataCollectionByAgentIdsInput, ...request.Option) (*applicationdiscoveryservice.StartDataCollectionByAgentIdsOutput, error)
	StartDataCollectionByAgentIdsRequest(*applicationdiscoveryservice.StartDataCollectionByAgentIdsInput) (*request.Request, *applicationdiscoveryservice.StartDataCollectionByAgentIdsOutput)

	StartExportTask(*applicationdiscoveryservice.StartExportTaskInput) (*applicationdiscoveryservice.StartExportTaskOutput, error)
	StartExportTaskWithContext(aws.Context, *applicationdiscoveryservice.StartExportTaskInput, ...request.Option) (*applicationdiscoveryservice.StartExportTaskOutput, error)
	StartExportTaskRequest(*applicationdiscoveryservice.StartExportTaskInput) (*request.Request, *applicationdiscoveryservice.StartExportTaskOutput)

	StartImportTask(*applicationdiscoveryservice.StartImportTaskInput) (*applicationdiscoveryservice.StartImportTaskOutput, error)
	StartImportTaskWithContext(aws.Context, *applicationdiscoveryservice.StartImportTaskInput, ...request.Option) (*applicationdiscoveryservice.StartImportTaskOutput, error)
	StartImportTaskRequest(*applicationdiscoveryservice.StartImportTaskInput) (*request.Request, *applicationdiscoveryservice.StartImportTaskOutput)

	StopContinuousExport(*applicationdiscoveryservice.StopContinuousExportInput) (*applicationdiscoveryservice.StopContinuousExportOutput, error)
	StopContinuousExportWithContext(aws.Context, *applicationdiscoveryservice.StopContinuousExportInput, ...request.Option) (*applicationdiscoveryservice.StopContinuousExportOutput, error)
	StopContinuousExportRequest(*applicationdiscoveryservice.StopContinuousExportInput) (*request.Request, *applicationdiscoveryservice.StopContinuousExportOutput)

	StopDataCollectionByAgentIds(*applicationdiscoveryservice.StopDataCollectionByAgentIdsInput) (*applicationdiscoveryservice.StopDataCollectionByAgentIdsOutput, error)
	StopDataCollectionByAgentIdsWithContext(aws.Context, *applicationdiscoveryservice.StopDataCollectionByAgentIdsInput, ...request.Option) (*applicationdiscoveryservice.StopDataCollectionByAgentIdsOutput, error)
	StopDataCollectionByAgentIdsRequest(*applicationdiscoveryservice.StopDataCollectionByAgentIdsInput) (*request.Request, *applicationdiscoveryservice.StopDataCollectionByAgentIdsOutput)

	UpdateApplication(*applicationdiscoveryservice.UpdateApplicationInput) (*applicationdiscoveryservice.UpdateApplicationOutput, error)
	UpdateApplicationWithContext(aws.Context, *applicationdiscoveryservice.UpdateApplicationInput, ...request.Option) (*applicationdiscoveryservice.UpdateApplicationOutput, error)
	UpdateApplicationRequest(*applicationdiscoveryservice.UpdateApplicationInput) (*request.Request, *applicationdiscoveryservice.UpdateApplicationOutput)
}

var _ ApplicationDiscoveryServiceAPI = (*applicationdiscoveryservice.ApplicationDiscoveryService)(nil)
