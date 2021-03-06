// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package applicationinsightsiface provides an interface to enable mocking the Amazon CloudWatch Application Insights service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package applicationinsightsiface

import (
	"github.com/pendo-io/aws-sdk-go/aws"
	"github.com/pendo-io/aws-sdk-go/aws/request"
	"github.com/pendo-io/aws-sdk-go/service/applicationinsights"
)

// ApplicationInsightsAPI provides an interface to enable mocking the
// applicationinsights.ApplicationInsights service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon CloudWatch Application Insights.
//    func myFunc(svc applicationinsightsiface.ApplicationInsightsAPI) bool {
//        // Make svc.CreateApplication request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := applicationinsights.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockApplicationInsightsClient struct {
//        applicationinsightsiface.ApplicationInsightsAPI
//    }
//    func (m *mockApplicationInsightsClient) CreateApplication(input *applicationinsights.CreateApplicationInput) (*applicationinsights.CreateApplicationOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockApplicationInsightsClient{}
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
type ApplicationInsightsAPI interface {
	CreateApplication(*applicationinsights.CreateApplicationInput) (*applicationinsights.CreateApplicationOutput, error)
	CreateApplicationWithContext(aws.Context, *applicationinsights.CreateApplicationInput, ...request.Option) (*applicationinsights.CreateApplicationOutput, error)
	CreateApplicationRequest(*applicationinsights.CreateApplicationInput) (*request.Request, *applicationinsights.CreateApplicationOutput)

	CreateComponent(*applicationinsights.CreateComponentInput) (*applicationinsights.CreateComponentOutput, error)
	CreateComponentWithContext(aws.Context, *applicationinsights.CreateComponentInput, ...request.Option) (*applicationinsights.CreateComponentOutput, error)
	CreateComponentRequest(*applicationinsights.CreateComponentInput) (*request.Request, *applicationinsights.CreateComponentOutput)

	DeleteApplication(*applicationinsights.DeleteApplicationInput) (*applicationinsights.DeleteApplicationOutput, error)
	DeleteApplicationWithContext(aws.Context, *applicationinsights.DeleteApplicationInput, ...request.Option) (*applicationinsights.DeleteApplicationOutput, error)
	DeleteApplicationRequest(*applicationinsights.DeleteApplicationInput) (*request.Request, *applicationinsights.DeleteApplicationOutput)

	DeleteComponent(*applicationinsights.DeleteComponentInput) (*applicationinsights.DeleteComponentOutput, error)
	DeleteComponentWithContext(aws.Context, *applicationinsights.DeleteComponentInput, ...request.Option) (*applicationinsights.DeleteComponentOutput, error)
	DeleteComponentRequest(*applicationinsights.DeleteComponentInput) (*request.Request, *applicationinsights.DeleteComponentOutput)

	DescribeApplication(*applicationinsights.DescribeApplicationInput) (*applicationinsights.DescribeApplicationOutput, error)
	DescribeApplicationWithContext(aws.Context, *applicationinsights.DescribeApplicationInput, ...request.Option) (*applicationinsights.DescribeApplicationOutput, error)
	DescribeApplicationRequest(*applicationinsights.DescribeApplicationInput) (*request.Request, *applicationinsights.DescribeApplicationOutput)

	DescribeComponent(*applicationinsights.DescribeComponentInput) (*applicationinsights.DescribeComponentOutput, error)
	DescribeComponentWithContext(aws.Context, *applicationinsights.DescribeComponentInput, ...request.Option) (*applicationinsights.DescribeComponentOutput, error)
	DescribeComponentRequest(*applicationinsights.DescribeComponentInput) (*request.Request, *applicationinsights.DescribeComponentOutput)

	DescribeComponentConfiguration(*applicationinsights.DescribeComponentConfigurationInput) (*applicationinsights.DescribeComponentConfigurationOutput, error)
	DescribeComponentConfigurationWithContext(aws.Context, *applicationinsights.DescribeComponentConfigurationInput, ...request.Option) (*applicationinsights.DescribeComponentConfigurationOutput, error)
	DescribeComponentConfigurationRequest(*applicationinsights.DescribeComponentConfigurationInput) (*request.Request, *applicationinsights.DescribeComponentConfigurationOutput)

	DescribeComponentConfigurationRecommendation(*applicationinsights.DescribeComponentConfigurationRecommendationInput) (*applicationinsights.DescribeComponentConfigurationRecommendationOutput, error)
	DescribeComponentConfigurationRecommendationWithContext(aws.Context, *applicationinsights.DescribeComponentConfigurationRecommendationInput, ...request.Option) (*applicationinsights.DescribeComponentConfigurationRecommendationOutput, error)
	DescribeComponentConfigurationRecommendationRequest(*applicationinsights.DescribeComponentConfigurationRecommendationInput) (*request.Request, *applicationinsights.DescribeComponentConfigurationRecommendationOutput)

	DescribeObservation(*applicationinsights.DescribeObservationInput) (*applicationinsights.DescribeObservationOutput, error)
	DescribeObservationWithContext(aws.Context, *applicationinsights.DescribeObservationInput, ...request.Option) (*applicationinsights.DescribeObservationOutput, error)
	DescribeObservationRequest(*applicationinsights.DescribeObservationInput) (*request.Request, *applicationinsights.DescribeObservationOutput)

	DescribeProblem(*applicationinsights.DescribeProblemInput) (*applicationinsights.DescribeProblemOutput, error)
	DescribeProblemWithContext(aws.Context, *applicationinsights.DescribeProblemInput, ...request.Option) (*applicationinsights.DescribeProblemOutput, error)
	DescribeProblemRequest(*applicationinsights.DescribeProblemInput) (*request.Request, *applicationinsights.DescribeProblemOutput)

	DescribeProblemObservations(*applicationinsights.DescribeProblemObservationsInput) (*applicationinsights.DescribeProblemObservationsOutput, error)
	DescribeProblemObservationsWithContext(aws.Context, *applicationinsights.DescribeProblemObservationsInput, ...request.Option) (*applicationinsights.DescribeProblemObservationsOutput, error)
	DescribeProblemObservationsRequest(*applicationinsights.DescribeProblemObservationsInput) (*request.Request, *applicationinsights.DescribeProblemObservationsOutput)

	ListApplications(*applicationinsights.ListApplicationsInput) (*applicationinsights.ListApplicationsOutput, error)
	ListApplicationsWithContext(aws.Context, *applicationinsights.ListApplicationsInput, ...request.Option) (*applicationinsights.ListApplicationsOutput, error)
	ListApplicationsRequest(*applicationinsights.ListApplicationsInput) (*request.Request, *applicationinsights.ListApplicationsOutput)

	ListApplicationsPages(*applicationinsights.ListApplicationsInput, func(*applicationinsights.ListApplicationsOutput, bool) bool) error
	ListApplicationsPagesWithContext(aws.Context, *applicationinsights.ListApplicationsInput, func(*applicationinsights.ListApplicationsOutput, bool) bool, ...request.Option) error

	ListComponents(*applicationinsights.ListComponentsInput) (*applicationinsights.ListComponentsOutput, error)
	ListComponentsWithContext(aws.Context, *applicationinsights.ListComponentsInput, ...request.Option) (*applicationinsights.ListComponentsOutput, error)
	ListComponentsRequest(*applicationinsights.ListComponentsInput) (*request.Request, *applicationinsights.ListComponentsOutput)

	ListComponentsPages(*applicationinsights.ListComponentsInput, func(*applicationinsights.ListComponentsOutput, bool) bool) error
	ListComponentsPagesWithContext(aws.Context, *applicationinsights.ListComponentsInput, func(*applicationinsights.ListComponentsOutput, bool) bool, ...request.Option) error

	ListProblems(*applicationinsights.ListProblemsInput) (*applicationinsights.ListProblemsOutput, error)
	ListProblemsWithContext(aws.Context, *applicationinsights.ListProblemsInput, ...request.Option) (*applicationinsights.ListProblemsOutput, error)
	ListProblemsRequest(*applicationinsights.ListProblemsInput) (*request.Request, *applicationinsights.ListProblemsOutput)

	ListProblemsPages(*applicationinsights.ListProblemsInput, func(*applicationinsights.ListProblemsOutput, bool) bool) error
	ListProblemsPagesWithContext(aws.Context, *applicationinsights.ListProblemsInput, func(*applicationinsights.ListProblemsOutput, bool) bool, ...request.Option) error

	UpdateComponent(*applicationinsights.UpdateComponentInput) (*applicationinsights.UpdateComponentOutput, error)
	UpdateComponentWithContext(aws.Context, *applicationinsights.UpdateComponentInput, ...request.Option) (*applicationinsights.UpdateComponentOutput, error)
	UpdateComponentRequest(*applicationinsights.UpdateComponentInput) (*request.Request, *applicationinsights.UpdateComponentOutput)

	UpdateComponentConfiguration(*applicationinsights.UpdateComponentConfigurationInput) (*applicationinsights.UpdateComponentConfigurationOutput, error)
	UpdateComponentConfigurationWithContext(aws.Context, *applicationinsights.UpdateComponentConfigurationInput, ...request.Option) (*applicationinsights.UpdateComponentConfigurationOutput, error)
	UpdateComponentConfigurationRequest(*applicationinsights.UpdateComponentConfigurationInput) (*request.Request, *applicationinsights.UpdateComponentConfigurationOutput)
}

var _ ApplicationInsightsAPI = (*applicationinsights.ApplicationInsights)(nil)
