// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package apigatewayv2iface provides an interface to enable mocking the AmazonApiGatewayV2 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package apigatewayv2iface

import (
	"github.com/pendo-io/aws-sdk-go/aws"
	"github.com/pendo-io/aws-sdk-go/aws/request"
	"github.com/pendo-io/aws-sdk-go/service/apigatewayv2"
)

// ApiGatewayV2API provides an interface to enable mocking the
// apigatewayv2.ApiGatewayV2 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AmazonApiGatewayV2.
//    func myFunc(svc apigatewayv2iface.ApiGatewayV2API) bool {
//        // Make svc.CreateApi request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := apigatewayv2.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockApiGatewayV2Client struct {
//        apigatewayv2iface.ApiGatewayV2API
//    }
//    func (m *mockApiGatewayV2Client) CreateApi(input *apigatewayv2.CreateApiInput) (*apigatewayv2.CreateApiOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockApiGatewayV2Client{}
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
type ApiGatewayV2API interface {
	CreateApi(*apigatewayv2.CreateApiInput) (*apigatewayv2.CreateApiOutput, error)
	CreateApiWithContext(aws.Context, *apigatewayv2.CreateApiInput, ...request.Option) (*apigatewayv2.CreateApiOutput, error)
	CreateApiRequest(*apigatewayv2.CreateApiInput) (*request.Request, *apigatewayv2.CreateApiOutput)

	CreateApiMapping(*apigatewayv2.CreateApiMappingInput) (*apigatewayv2.CreateApiMappingOutput, error)
	CreateApiMappingWithContext(aws.Context, *apigatewayv2.CreateApiMappingInput, ...request.Option) (*apigatewayv2.CreateApiMappingOutput, error)
	CreateApiMappingRequest(*apigatewayv2.CreateApiMappingInput) (*request.Request, *apigatewayv2.CreateApiMappingOutput)

	CreateAuthorizer(*apigatewayv2.CreateAuthorizerInput) (*apigatewayv2.CreateAuthorizerOutput, error)
	CreateAuthorizerWithContext(aws.Context, *apigatewayv2.CreateAuthorizerInput, ...request.Option) (*apigatewayv2.CreateAuthorizerOutput, error)
	CreateAuthorizerRequest(*apigatewayv2.CreateAuthorizerInput) (*request.Request, *apigatewayv2.CreateAuthorizerOutput)

	CreateDeployment(*apigatewayv2.CreateDeploymentInput) (*apigatewayv2.CreateDeploymentOutput, error)
	CreateDeploymentWithContext(aws.Context, *apigatewayv2.CreateDeploymentInput, ...request.Option) (*apigatewayv2.CreateDeploymentOutput, error)
	CreateDeploymentRequest(*apigatewayv2.CreateDeploymentInput) (*request.Request, *apigatewayv2.CreateDeploymentOutput)

	CreateDomainName(*apigatewayv2.CreateDomainNameInput) (*apigatewayv2.CreateDomainNameOutput, error)
	CreateDomainNameWithContext(aws.Context, *apigatewayv2.CreateDomainNameInput, ...request.Option) (*apigatewayv2.CreateDomainNameOutput, error)
	CreateDomainNameRequest(*apigatewayv2.CreateDomainNameInput) (*request.Request, *apigatewayv2.CreateDomainNameOutput)

	CreateIntegration(*apigatewayv2.CreateIntegrationInput) (*apigatewayv2.CreateIntegrationOutput, error)
	CreateIntegrationWithContext(aws.Context, *apigatewayv2.CreateIntegrationInput, ...request.Option) (*apigatewayv2.CreateIntegrationOutput, error)
	CreateIntegrationRequest(*apigatewayv2.CreateIntegrationInput) (*request.Request, *apigatewayv2.CreateIntegrationOutput)

	CreateIntegrationResponse(*apigatewayv2.CreateIntegrationResponseInput) (*apigatewayv2.CreateIntegrationResponseOutput, error)
	CreateIntegrationResponseWithContext(aws.Context, *apigatewayv2.CreateIntegrationResponseInput, ...request.Option) (*apigatewayv2.CreateIntegrationResponseOutput, error)
	CreateIntegrationResponseRequest(*apigatewayv2.CreateIntegrationResponseInput) (*request.Request, *apigatewayv2.CreateIntegrationResponseOutput)

	CreateModel(*apigatewayv2.CreateModelInput) (*apigatewayv2.CreateModelOutput, error)
	CreateModelWithContext(aws.Context, *apigatewayv2.CreateModelInput, ...request.Option) (*apigatewayv2.CreateModelOutput, error)
	CreateModelRequest(*apigatewayv2.CreateModelInput) (*request.Request, *apigatewayv2.CreateModelOutput)

	CreateRoute(*apigatewayv2.CreateRouteInput) (*apigatewayv2.CreateRouteOutput, error)
	CreateRouteWithContext(aws.Context, *apigatewayv2.CreateRouteInput, ...request.Option) (*apigatewayv2.CreateRouteOutput, error)
	CreateRouteRequest(*apigatewayv2.CreateRouteInput) (*request.Request, *apigatewayv2.CreateRouteOutput)

	CreateRouteResponse(*apigatewayv2.CreateRouteResponseInput) (*apigatewayv2.CreateRouteResponseOutput, error)
	CreateRouteResponseWithContext(aws.Context, *apigatewayv2.CreateRouteResponseInput, ...request.Option) (*apigatewayv2.CreateRouteResponseOutput, error)
	CreateRouteResponseRequest(*apigatewayv2.CreateRouteResponseInput) (*request.Request, *apigatewayv2.CreateRouteResponseOutput)

	CreateStage(*apigatewayv2.CreateStageInput) (*apigatewayv2.CreateStageOutput, error)
	CreateStageWithContext(aws.Context, *apigatewayv2.CreateStageInput, ...request.Option) (*apigatewayv2.CreateStageOutput, error)
	CreateStageRequest(*apigatewayv2.CreateStageInput) (*request.Request, *apigatewayv2.CreateStageOutput)

	DeleteApi(*apigatewayv2.DeleteApiInput) (*apigatewayv2.DeleteApiOutput, error)
	DeleteApiWithContext(aws.Context, *apigatewayv2.DeleteApiInput, ...request.Option) (*apigatewayv2.DeleteApiOutput, error)
	DeleteApiRequest(*apigatewayv2.DeleteApiInput) (*request.Request, *apigatewayv2.DeleteApiOutput)

	DeleteApiMapping(*apigatewayv2.DeleteApiMappingInput) (*apigatewayv2.DeleteApiMappingOutput, error)
	DeleteApiMappingWithContext(aws.Context, *apigatewayv2.DeleteApiMappingInput, ...request.Option) (*apigatewayv2.DeleteApiMappingOutput, error)
	DeleteApiMappingRequest(*apigatewayv2.DeleteApiMappingInput) (*request.Request, *apigatewayv2.DeleteApiMappingOutput)

	DeleteAuthorizer(*apigatewayv2.DeleteAuthorizerInput) (*apigatewayv2.DeleteAuthorizerOutput, error)
	DeleteAuthorizerWithContext(aws.Context, *apigatewayv2.DeleteAuthorizerInput, ...request.Option) (*apigatewayv2.DeleteAuthorizerOutput, error)
	DeleteAuthorizerRequest(*apigatewayv2.DeleteAuthorizerInput) (*request.Request, *apigatewayv2.DeleteAuthorizerOutput)

	DeleteDeployment(*apigatewayv2.DeleteDeploymentInput) (*apigatewayv2.DeleteDeploymentOutput, error)
	DeleteDeploymentWithContext(aws.Context, *apigatewayv2.DeleteDeploymentInput, ...request.Option) (*apigatewayv2.DeleteDeploymentOutput, error)
	DeleteDeploymentRequest(*apigatewayv2.DeleteDeploymentInput) (*request.Request, *apigatewayv2.DeleteDeploymentOutput)

	DeleteDomainName(*apigatewayv2.DeleteDomainNameInput) (*apigatewayv2.DeleteDomainNameOutput, error)
	DeleteDomainNameWithContext(aws.Context, *apigatewayv2.DeleteDomainNameInput, ...request.Option) (*apigatewayv2.DeleteDomainNameOutput, error)
	DeleteDomainNameRequest(*apigatewayv2.DeleteDomainNameInput) (*request.Request, *apigatewayv2.DeleteDomainNameOutput)

	DeleteIntegration(*apigatewayv2.DeleteIntegrationInput) (*apigatewayv2.DeleteIntegrationOutput, error)
	DeleteIntegrationWithContext(aws.Context, *apigatewayv2.DeleteIntegrationInput, ...request.Option) (*apigatewayv2.DeleteIntegrationOutput, error)
	DeleteIntegrationRequest(*apigatewayv2.DeleteIntegrationInput) (*request.Request, *apigatewayv2.DeleteIntegrationOutput)

	DeleteIntegrationResponse(*apigatewayv2.DeleteIntegrationResponseInput) (*apigatewayv2.DeleteIntegrationResponseOutput, error)
	DeleteIntegrationResponseWithContext(aws.Context, *apigatewayv2.DeleteIntegrationResponseInput, ...request.Option) (*apigatewayv2.DeleteIntegrationResponseOutput, error)
	DeleteIntegrationResponseRequest(*apigatewayv2.DeleteIntegrationResponseInput) (*request.Request, *apigatewayv2.DeleteIntegrationResponseOutput)

	DeleteModel(*apigatewayv2.DeleteModelInput) (*apigatewayv2.DeleteModelOutput, error)
	DeleteModelWithContext(aws.Context, *apigatewayv2.DeleteModelInput, ...request.Option) (*apigatewayv2.DeleteModelOutput, error)
	DeleteModelRequest(*apigatewayv2.DeleteModelInput) (*request.Request, *apigatewayv2.DeleteModelOutput)

	DeleteRoute(*apigatewayv2.DeleteRouteInput) (*apigatewayv2.DeleteRouteOutput, error)
	DeleteRouteWithContext(aws.Context, *apigatewayv2.DeleteRouteInput, ...request.Option) (*apigatewayv2.DeleteRouteOutput, error)
	DeleteRouteRequest(*apigatewayv2.DeleteRouteInput) (*request.Request, *apigatewayv2.DeleteRouteOutput)

	DeleteRouteResponse(*apigatewayv2.DeleteRouteResponseInput) (*apigatewayv2.DeleteRouteResponseOutput, error)
	DeleteRouteResponseWithContext(aws.Context, *apigatewayv2.DeleteRouteResponseInput, ...request.Option) (*apigatewayv2.DeleteRouteResponseOutput, error)
	DeleteRouteResponseRequest(*apigatewayv2.DeleteRouteResponseInput) (*request.Request, *apigatewayv2.DeleteRouteResponseOutput)

	DeleteStage(*apigatewayv2.DeleteStageInput) (*apigatewayv2.DeleteStageOutput, error)
	DeleteStageWithContext(aws.Context, *apigatewayv2.DeleteStageInput, ...request.Option) (*apigatewayv2.DeleteStageOutput, error)
	DeleteStageRequest(*apigatewayv2.DeleteStageInput) (*request.Request, *apigatewayv2.DeleteStageOutput)

	GetApi(*apigatewayv2.GetApiInput) (*apigatewayv2.GetApiOutput, error)
	GetApiWithContext(aws.Context, *apigatewayv2.GetApiInput, ...request.Option) (*apigatewayv2.GetApiOutput, error)
	GetApiRequest(*apigatewayv2.GetApiInput) (*request.Request, *apigatewayv2.GetApiOutput)

	GetApiMapping(*apigatewayv2.GetApiMappingInput) (*apigatewayv2.GetApiMappingOutput, error)
	GetApiMappingWithContext(aws.Context, *apigatewayv2.GetApiMappingInput, ...request.Option) (*apigatewayv2.GetApiMappingOutput, error)
	GetApiMappingRequest(*apigatewayv2.GetApiMappingInput) (*request.Request, *apigatewayv2.GetApiMappingOutput)

	GetApiMappings(*apigatewayv2.GetApiMappingsInput) (*apigatewayv2.GetApiMappingsOutput, error)
	GetApiMappingsWithContext(aws.Context, *apigatewayv2.GetApiMappingsInput, ...request.Option) (*apigatewayv2.GetApiMappingsOutput, error)
	GetApiMappingsRequest(*apigatewayv2.GetApiMappingsInput) (*request.Request, *apigatewayv2.GetApiMappingsOutput)

	GetApis(*apigatewayv2.GetApisInput) (*apigatewayv2.GetApisOutput, error)
	GetApisWithContext(aws.Context, *apigatewayv2.GetApisInput, ...request.Option) (*apigatewayv2.GetApisOutput, error)
	GetApisRequest(*apigatewayv2.GetApisInput) (*request.Request, *apigatewayv2.GetApisOutput)

	GetAuthorizer(*apigatewayv2.GetAuthorizerInput) (*apigatewayv2.GetAuthorizerOutput, error)
	GetAuthorizerWithContext(aws.Context, *apigatewayv2.GetAuthorizerInput, ...request.Option) (*apigatewayv2.GetAuthorizerOutput, error)
	GetAuthorizerRequest(*apigatewayv2.GetAuthorizerInput) (*request.Request, *apigatewayv2.GetAuthorizerOutput)

	GetAuthorizers(*apigatewayv2.GetAuthorizersInput) (*apigatewayv2.GetAuthorizersOutput, error)
	GetAuthorizersWithContext(aws.Context, *apigatewayv2.GetAuthorizersInput, ...request.Option) (*apigatewayv2.GetAuthorizersOutput, error)
	GetAuthorizersRequest(*apigatewayv2.GetAuthorizersInput) (*request.Request, *apigatewayv2.GetAuthorizersOutput)

	GetDeployment(*apigatewayv2.GetDeploymentInput) (*apigatewayv2.GetDeploymentOutput, error)
	GetDeploymentWithContext(aws.Context, *apigatewayv2.GetDeploymentInput, ...request.Option) (*apigatewayv2.GetDeploymentOutput, error)
	GetDeploymentRequest(*apigatewayv2.GetDeploymentInput) (*request.Request, *apigatewayv2.GetDeploymentOutput)

	GetDeployments(*apigatewayv2.GetDeploymentsInput) (*apigatewayv2.GetDeploymentsOutput, error)
	GetDeploymentsWithContext(aws.Context, *apigatewayv2.GetDeploymentsInput, ...request.Option) (*apigatewayv2.GetDeploymentsOutput, error)
	GetDeploymentsRequest(*apigatewayv2.GetDeploymentsInput) (*request.Request, *apigatewayv2.GetDeploymentsOutput)

	GetDomainName(*apigatewayv2.GetDomainNameInput) (*apigatewayv2.GetDomainNameOutput, error)
	GetDomainNameWithContext(aws.Context, *apigatewayv2.GetDomainNameInput, ...request.Option) (*apigatewayv2.GetDomainNameOutput, error)
	GetDomainNameRequest(*apigatewayv2.GetDomainNameInput) (*request.Request, *apigatewayv2.GetDomainNameOutput)

	GetDomainNames(*apigatewayv2.GetDomainNamesInput) (*apigatewayv2.GetDomainNamesOutput, error)
	GetDomainNamesWithContext(aws.Context, *apigatewayv2.GetDomainNamesInput, ...request.Option) (*apigatewayv2.GetDomainNamesOutput, error)
	GetDomainNamesRequest(*apigatewayv2.GetDomainNamesInput) (*request.Request, *apigatewayv2.GetDomainNamesOutput)

	GetIntegration(*apigatewayv2.GetIntegrationInput) (*apigatewayv2.GetIntegrationOutput, error)
	GetIntegrationWithContext(aws.Context, *apigatewayv2.GetIntegrationInput, ...request.Option) (*apigatewayv2.GetIntegrationOutput, error)
	GetIntegrationRequest(*apigatewayv2.GetIntegrationInput) (*request.Request, *apigatewayv2.GetIntegrationOutput)

	GetIntegrationResponse(*apigatewayv2.GetIntegrationResponseInput) (*apigatewayv2.GetIntegrationResponseOutput, error)
	GetIntegrationResponseWithContext(aws.Context, *apigatewayv2.GetIntegrationResponseInput, ...request.Option) (*apigatewayv2.GetIntegrationResponseOutput, error)
	GetIntegrationResponseRequest(*apigatewayv2.GetIntegrationResponseInput) (*request.Request, *apigatewayv2.GetIntegrationResponseOutput)

	GetIntegrationResponses(*apigatewayv2.GetIntegrationResponsesInput) (*apigatewayv2.GetIntegrationResponsesOutput, error)
	GetIntegrationResponsesWithContext(aws.Context, *apigatewayv2.GetIntegrationResponsesInput, ...request.Option) (*apigatewayv2.GetIntegrationResponsesOutput, error)
	GetIntegrationResponsesRequest(*apigatewayv2.GetIntegrationResponsesInput) (*request.Request, *apigatewayv2.GetIntegrationResponsesOutput)

	GetIntegrations(*apigatewayv2.GetIntegrationsInput) (*apigatewayv2.GetIntegrationsOutput, error)
	GetIntegrationsWithContext(aws.Context, *apigatewayv2.GetIntegrationsInput, ...request.Option) (*apigatewayv2.GetIntegrationsOutput, error)
	GetIntegrationsRequest(*apigatewayv2.GetIntegrationsInput) (*request.Request, *apigatewayv2.GetIntegrationsOutput)

	GetModel(*apigatewayv2.GetModelInput) (*apigatewayv2.GetModelOutput, error)
	GetModelWithContext(aws.Context, *apigatewayv2.GetModelInput, ...request.Option) (*apigatewayv2.GetModelOutput, error)
	GetModelRequest(*apigatewayv2.GetModelInput) (*request.Request, *apigatewayv2.GetModelOutput)

	GetModelTemplate(*apigatewayv2.GetModelTemplateInput) (*apigatewayv2.GetModelTemplateOutput, error)
	GetModelTemplateWithContext(aws.Context, *apigatewayv2.GetModelTemplateInput, ...request.Option) (*apigatewayv2.GetModelTemplateOutput, error)
	GetModelTemplateRequest(*apigatewayv2.GetModelTemplateInput) (*request.Request, *apigatewayv2.GetModelTemplateOutput)

	GetModels(*apigatewayv2.GetModelsInput) (*apigatewayv2.GetModelsOutput, error)
	GetModelsWithContext(aws.Context, *apigatewayv2.GetModelsInput, ...request.Option) (*apigatewayv2.GetModelsOutput, error)
	GetModelsRequest(*apigatewayv2.GetModelsInput) (*request.Request, *apigatewayv2.GetModelsOutput)

	GetRoute(*apigatewayv2.GetRouteInput) (*apigatewayv2.GetRouteOutput, error)
	GetRouteWithContext(aws.Context, *apigatewayv2.GetRouteInput, ...request.Option) (*apigatewayv2.GetRouteOutput, error)
	GetRouteRequest(*apigatewayv2.GetRouteInput) (*request.Request, *apigatewayv2.GetRouteOutput)

	GetRouteResponse(*apigatewayv2.GetRouteResponseInput) (*apigatewayv2.GetRouteResponseOutput, error)
	GetRouteResponseWithContext(aws.Context, *apigatewayv2.GetRouteResponseInput, ...request.Option) (*apigatewayv2.GetRouteResponseOutput, error)
	GetRouteResponseRequest(*apigatewayv2.GetRouteResponseInput) (*request.Request, *apigatewayv2.GetRouteResponseOutput)

	GetRouteResponses(*apigatewayv2.GetRouteResponsesInput) (*apigatewayv2.GetRouteResponsesOutput, error)
	GetRouteResponsesWithContext(aws.Context, *apigatewayv2.GetRouteResponsesInput, ...request.Option) (*apigatewayv2.GetRouteResponsesOutput, error)
	GetRouteResponsesRequest(*apigatewayv2.GetRouteResponsesInput) (*request.Request, *apigatewayv2.GetRouteResponsesOutput)

	GetRoutes(*apigatewayv2.GetRoutesInput) (*apigatewayv2.GetRoutesOutput, error)
	GetRoutesWithContext(aws.Context, *apigatewayv2.GetRoutesInput, ...request.Option) (*apigatewayv2.GetRoutesOutput, error)
	GetRoutesRequest(*apigatewayv2.GetRoutesInput) (*request.Request, *apigatewayv2.GetRoutesOutput)

	GetStage(*apigatewayv2.GetStageInput) (*apigatewayv2.GetStageOutput, error)
	GetStageWithContext(aws.Context, *apigatewayv2.GetStageInput, ...request.Option) (*apigatewayv2.GetStageOutput, error)
	GetStageRequest(*apigatewayv2.GetStageInput) (*request.Request, *apigatewayv2.GetStageOutput)

	GetStages(*apigatewayv2.GetStagesInput) (*apigatewayv2.GetStagesOutput, error)
	GetStagesWithContext(aws.Context, *apigatewayv2.GetStagesInput, ...request.Option) (*apigatewayv2.GetStagesOutput, error)
	GetStagesRequest(*apigatewayv2.GetStagesInput) (*request.Request, *apigatewayv2.GetStagesOutput)

	GetTags(*apigatewayv2.GetTagsInput) (*apigatewayv2.GetTagsOutput, error)
	GetTagsWithContext(aws.Context, *apigatewayv2.GetTagsInput, ...request.Option) (*apigatewayv2.GetTagsOutput, error)
	GetTagsRequest(*apigatewayv2.GetTagsInput) (*request.Request, *apigatewayv2.GetTagsOutput)

	TagResource(*apigatewayv2.TagResourceInput) (*apigatewayv2.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *apigatewayv2.TagResourceInput, ...request.Option) (*apigatewayv2.TagResourceOutput, error)
	TagResourceRequest(*apigatewayv2.TagResourceInput) (*request.Request, *apigatewayv2.TagResourceOutput)

	UntagResource(*apigatewayv2.UntagResourceInput) (*apigatewayv2.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *apigatewayv2.UntagResourceInput, ...request.Option) (*apigatewayv2.UntagResourceOutput, error)
	UntagResourceRequest(*apigatewayv2.UntagResourceInput) (*request.Request, *apigatewayv2.UntagResourceOutput)

	UpdateApi(*apigatewayv2.UpdateApiInput) (*apigatewayv2.UpdateApiOutput, error)
	UpdateApiWithContext(aws.Context, *apigatewayv2.UpdateApiInput, ...request.Option) (*apigatewayv2.UpdateApiOutput, error)
	UpdateApiRequest(*apigatewayv2.UpdateApiInput) (*request.Request, *apigatewayv2.UpdateApiOutput)

	UpdateApiMapping(*apigatewayv2.UpdateApiMappingInput) (*apigatewayv2.UpdateApiMappingOutput, error)
	UpdateApiMappingWithContext(aws.Context, *apigatewayv2.UpdateApiMappingInput, ...request.Option) (*apigatewayv2.UpdateApiMappingOutput, error)
	UpdateApiMappingRequest(*apigatewayv2.UpdateApiMappingInput) (*request.Request, *apigatewayv2.UpdateApiMappingOutput)

	UpdateAuthorizer(*apigatewayv2.UpdateAuthorizerInput) (*apigatewayv2.UpdateAuthorizerOutput, error)
	UpdateAuthorizerWithContext(aws.Context, *apigatewayv2.UpdateAuthorizerInput, ...request.Option) (*apigatewayv2.UpdateAuthorizerOutput, error)
	UpdateAuthorizerRequest(*apigatewayv2.UpdateAuthorizerInput) (*request.Request, *apigatewayv2.UpdateAuthorizerOutput)

	UpdateDeployment(*apigatewayv2.UpdateDeploymentInput) (*apigatewayv2.UpdateDeploymentOutput, error)
	UpdateDeploymentWithContext(aws.Context, *apigatewayv2.UpdateDeploymentInput, ...request.Option) (*apigatewayv2.UpdateDeploymentOutput, error)
	UpdateDeploymentRequest(*apigatewayv2.UpdateDeploymentInput) (*request.Request, *apigatewayv2.UpdateDeploymentOutput)

	UpdateDomainName(*apigatewayv2.UpdateDomainNameInput) (*apigatewayv2.UpdateDomainNameOutput, error)
	UpdateDomainNameWithContext(aws.Context, *apigatewayv2.UpdateDomainNameInput, ...request.Option) (*apigatewayv2.UpdateDomainNameOutput, error)
	UpdateDomainNameRequest(*apigatewayv2.UpdateDomainNameInput) (*request.Request, *apigatewayv2.UpdateDomainNameOutput)

	UpdateIntegration(*apigatewayv2.UpdateIntegrationInput) (*apigatewayv2.UpdateIntegrationOutput, error)
	UpdateIntegrationWithContext(aws.Context, *apigatewayv2.UpdateIntegrationInput, ...request.Option) (*apigatewayv2.UpdateIntegrationOutput, error)
	UpdateIntegrationRequest(*apigatewayv2.UpdateIntegrationInput) (*request.Request, *apigatewayv2.UpdateIntegrationOutput)

	UpdateIntegrationResponse(*apigatewayv2.UpdateIntegrationResponseInput) (*apigatewayv2.UpdateIntegrationResponseOutput, error)
	UpdateIntegrationResponseWithContext(aws.Context, *apigatewayv2.UpdateIntegrationResponseInput, ...request.Option) (*apigatewayv2.UpdateIntegrationResponseOutput, error)
	UpdateIntegrationResponseRequest(*apigatewayv2.UpdateIntegrationResponseInput) (*request.Request, *apigatewayv2.UpdateIntegrationResponseOutput)

	UpdateModel(*apigatewayv2.UpdateModelInput) (*apigatewayv2.UpdateModelOutput, error)
	UpdateModelWithContext(aws.Context, *apigatewayv2.UpdateModelInput, ...request.Option) (*apigatewayv2.UpdateModelOutput, error)
	UpdateModelRequest(*apigatewayv2.UpdateModelInput) (*request.Request, *apigatewayv2.UpdateModelOutput)

	UpdateRoute(*apigatewayv2.UpdateRouteInput) (*apigatewayv2.UpdateRouteOutput, error)
	UpdateRouteWithContext(aws.Context, *apigatewayv2.UpdateRouteInput, ...request.Option) (*apigatewayv2.UpdateRouteOutput, error)
	UpdateRouteRequest(*apigatewayv2.UpdateRouteInput) (*request.Request, *apigatewayv2.UpdateRouteOutput)

	UpdateRouteResponse(*apigatewayv2.UpdateRouteResponseInput) (*apigatewayv2.UpdateRouteResponseOutput, error)
	UpdateRouteResponseWithContext(aws.Context, *apigatewayv2.UpdateRouteResponseInput, ...request.Option) (*apigatewayv2.UpdateRouteResponseOutput, error)
	UpdateRouteResponseRequest(*apigatewayv2.UpdateRouteResponseInput) (*request.Request, *apigatewayv2.UpdateRouteResponseOutput)

	UpdateStage(*apigatewayv2.UpdateStageInput) (*apigatewayv2.UpdateStageOutput, error)
	UpdateStageWithContext(aws.Context, *apigatewayv2.UpdateStageInput, ...request.Option) (*apigatewayv2.UpdateStageOutput, error)
	UpdateStageRequest(*apigatewayv2.UpdateStageInput) (*request.Request, *apigatewayv2.UpdateStageOutput)
}

var _ ApiGatewayV2API = (*apigatewayv2.ApiGatewayV2)(nil)
