// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package servicediscoveryiface provides an interface to enable mocking the AWS Cloud Map service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package servicediscoveryiface

import (
	"github.com/pendo-io/aws-sdk-go/aws"
	"github.com/pendo-io/aws-sdk-go/aws/request"
	"github.com/pendo-io/aws-sdk-go/service/servicediscovery"
)

// ServiceDiscoveryAPI provides an interface to enable mocking the
// servicediscovery.ServiceDiscovery service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Cloud Map.
//    func myFunc(svc servicediscoveryiface.ServiceDiscoveryAPI) bool {
//        // Make svc.CreateHttpNamespace request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := servicediscovery.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockServiceDiscoveryClient struct {
//        servicediscoveryiface.ServiceDiscoveryAPI
//    }
//    func (m *mockServiceDiscoveryClient) CreateHttpNamespace(input *servicediscovery.CreateHttpNamespaceInput) (*servicediscovery.CreateHttpNamespaceOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockServiceDiscoveryClient{}
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
type ServiceDiscoveryAPI interface {
	CreateHttpNamespace(*servicediscovery.CreateHttpNamespaceInput) (*servicediscovery.CreateHttpNamespaceOutput, error)
	CreateHttpNamespaceWithContext(aws.Context, *servicediscovery.CreateHttpNamespaceInput, ...request.Option) (*servicediscovery.CreateHttpNamespaceOutput, error)
	CreateHttpNamespaceRequest(*servicediscovery.CreateHttpNamespaceInput) (*request.Request, *servicediscovery.CreateHttpNamespaceOutput)

	CreatePrivateDnsNamespace(*servicediscovery.CreatePrivateDnsNamespaceInput) (*servicediscovery.CreatePrivateDnsNamespaceOutput, error)
	CreatePrivateDnsNamespaceWithContext(aws.Context, *servicediscovery.CreatePrivateDnsNamespaceInput, ...request.Option) (*servicediscovery.CreatePrivateDnsNamespaceOutput, error)
	CreatePrivateDnsNamespaceRequest(*servicediscovery.CreatePrivateDnsNamespaceInput) (*request.Request, *servicediscovery.CreatePrivateDnsNamespaceOutput)

	CreatePublicDnsNamespace(*servicediscovery.CreatePublicDnsNamespaceInput) (*servicediscovery.CreatePublicDnsNamespaceOutput, error)
	CreatePublicDnsNamespaceWithContext(aws.Context, *servicediscovery.CreatePublicDnsNamespaceInput, ...request.Option) (*servicediscovery.CreatePublicDnsNamespaceOutput, error)
	CreatePublicDnsNamespaceRequest(*servicediscovery.CreatePublicDnsNamespaceInput) (*request.Request, *servicediscovery.CreatePublicDnsNamespaceOutput)

	CreateService(*servicediscovery.CreateServiceInput) (*servicediscovery.CreateServiceOutput, error)
	CreateServiceWithContext(aws.Context, *servicediscovery.CreateServiceInput, ...request.Option) (*servicediscovery.CreateServiceOutput, error)
	CreateServiceRequest(*servicediscovery.CreateServiceInput) (*request.Request, *servicediscovery.CreateServiceOutput)

	DeleteNamespace(*servicediscovery.DeleteNamespaceInput) (*servicediscovery.DeleteNamespaceOutput, error)
	DeleteNamespaceWithContext(aws.Context, *servicediscovery.DeleteNamespaceInput, ...request.Option) (*servicediscovery.DeleteNamespaceOutput, error)
	DeleteNamespaceRequest(*servicediscovery.DeleteNamespaceInput) (*request.Request, *servicediscovery.DeleteNamespaceOutput)

	DeleteService(*servicediscovery.DeleteServiceInput) (*servicediscovery.DeleteServiceOutput, error)
	DeleteServiceWithContext(aws.Context, *servicediscovery.DeleteServiceInput, ...request.Option) (*servicediscovery.DeleteServiceOutput, error)
	DeleteServiceRequest(*servicediscovery.DeleteServiceInput) (*request.Request, *servicediscovery.DeleteServiceOutput)

	DeregisterInstance(*servicediscovery.DeregisterInstanceInput) (*servicediscovery.DeregisterInstanceOutput, error)
	DeregisterInstanceWithContext(aws.Context, *servicediscovery.DeregisterInstanceInput, ...request.Option) (*servicediscovery.DeregisterInstanceOutput, error)
	DeregisterInstanceRequest(*servicediscovery.DeregisterInstanceInput) (*request.Request, *servicediscovery.DeregisterInstanceOutput)

	DiscoverInstances(*servicediscovery.DiscoverInstancesInput) (*servicediscovery.DiscoverInstancesOutput, error)
	DiscoverInstancesWithContext(aws.Context, *servicediscovery.DiscoverInstancesInput, ...request.Option) (*servicediscovery.DiscoverInstancesOutput, error)
	DiscoverInstancesRequest(*servicediscovery.DiscoverInstancesInput) (*request.Request, *servicediscovery.DiscoverInstancesOutput)

	GetInstance(*servicediscovery.GetInstanceInput) (*servicediscovery.GetInstanceOutput, error)
	GetInstanceWithContext(aws.Context, *servicediscovery.GetInstanceInput, ...request.Option) (*servicediscovery.GetInstanceOutput, error)
	GetInstanceRequest(*servicediscovery.GetInstanceInput) (*request.Request, *servicediscovery.GetInstanceOutput)

	GetInstancesHealthStatus(*servicediscovery.GetInstancesHealthStatusInput) (*servicediscovery.GetInstancesHealthStatusOutput, error)
	GetInstancesHealthStatusWithContext(aws.Context, *servicediscovery.GetInstancesHealthStatusInput, ...request.Option) (*servicediscovery.GetInstancesHealthStatusOutput, error)
	GetInstancesHealthStatusRequest(*servicediscovery.GetInstancesHealthStatusInput) (*request.Request, *servicediscovery.GetInstancesHealthStatusOutput)

	GetInstancesHealthStatusPages(*servicediscovery.GetInstancesHealthStatusInput, func(*servicediscovery.GetInstancesHealthStatusOutput, bool) bool) error
	GetInstancesHealthStatusPagesWithContext(aws.Context, *servicediscovery.GetInstancesHealthStatusInput, func(*servicediscovery.GetInstancesHealthStatusOutput, bool) bool, ...request.Option) error

	GetNamespace(*servicediscovery.GetNamespaceInput) (*servicediscovery.GetNamespaceOutput, error)
	GetNamespaceWithContext(aws.Context, *servicediscovery.GetNamespaceInput, ...request.Option) (*servicediscovery.GetNamespaceOutput, error)
	GetNamespaceRequest(*servicediscovery.GetNamespaceInput) (*request.Request, *servicediscovery.GetNamespaceOutput)

	GetOperation(*servicediscovery.GetOperationInput) (*servicediscovery.GetOperationOutput, error)
	GetOperationWithContext(aws.Context, *servicediscovery.GetOperationInput, ...request.Option) (*servicediscovery.GetOperationOutput, error)
	GetOperationRequest(*servicediscovery.GetOperationInput) (*request.Request, *servicediscovery.GetOperationOutput)

	GetService(*servicediscovery.GetServiceInput) (*servicediscovery.GetServiceOutput, error)
	GetServiceWithContext(aws.Context, *servicediscovery.GetServiceInput, ...request.Option) (*servicediscovery.GetServiceOutput, error)
	GetServiceRequest(*servicediscovery.GetServiceInput) (*request.Request, *servicediscovery.GetServiceOutput)

	ListInstances(*servicediscovery.ListInstancesInput) (*servicediscovery.ListInstancesOutput, error)
	ListInstancesWithContext(aws.Context, *servicediscovery.ListInstancesInput, ...request.Option) (*servicediscovery.ListInstancesOutput, error)
	ListInstancesRequest(*servicediscovery.ListInstancesInput) (*request.Request, *servicediscovery.ListInstancesOutput)

	ListInstancesPages(*servicediscovery.ListInstancesInput, func(*servicediscovery.ListInstancesOutput, bool) bool) error
	ListInstancesPagesWithContext(aws.Context, *servicediscovery.ListInstancesInput, func(*servicediscovery.ListInstancesOutput, bool) bool, ...request.Option) error

	ListNamespaces(*servicediscovery.ListNamespacesInput) (*servicediscovery.ListNamespacesOutput, error)
	ListNamespacesWithContext(aws.Context, *servicediscovery.ListNamespacesInput, ...request.Option) (*servicediscovery.ListNamespacesOutput, error)
	ListNamespacesRequest(*servicediscovery.ListNamespacesInput) (*request.Request, *servicediscovery.ListNamespacesOutput)

	ListNamespacesPages(*servicediscovery.ListNamespacesInput, func(*servicediscovery.ListNamespacesOutput, bool) bool) error
	ListNamespacesPagesWithContext(aws.Context, *servicediscovery.ListNamespacesInput, func(*servicediscovery.ListNamespacesOutput, bool) bool, ...request.Option) error

	ListOperations(*servicediscovery.ListOperationsInput) (*servicediscovery.ListOperationsOutput, error)
	ListOperationsWithContext(aws.Context, *servicediscovery.ListOperationsInput, ...request.Option) (*servicediscovery.ListOperationsOutput, error)
	ListOperationsRequest(*servicediscovery.ListOperationsInput) (*request.Request, *servicediscovery.ListOperationsOutput)

	ListOperationsPages(*servicediscovery.ListOperationsInput, func(*servicediscovery.ListOperationsOutput, bool) bool) error
	ListOperationsPagesWithContext(aws.Context, *servicediscovery.ListOperationsInput, func(*servicediscovery.ListOperationsOutput, bool) bool, ...request.Option) error

	ListServices(*servicediscovery.ListServicesInput) (*servicediscovery.ListServicesOutput, error)
	ListServicesWithContext(aws.Context, *servicediscovery.ListServicesInput, ...request.Option) (*servicediscovery.ListServicesOutput, error)
	ListServicesRequest(*servicediscovery.ListServicesInput) (*request.Request, *servicediscovery.ListServicesOutput)

	ListServicesPages(*servicediscovery.ListServicesInput, func(*servicediscovery.ListServicesOutput, bool) bool) error
	ListServicesPagesWithContext(aws.Context, *servicediscovery.ListServicesInput, func(*servicediscovery.ListServicesOutput, bool) bool, ...request.Option) error

	RegisterInstance(*servicediscovery.RegisterInstanceInput) (*servicediscovery.RegisterInstanceOutput, error)
	RegisterInstanceWithContext(aws.Context, *servicediscovery.RegisterInstanceInput, ...request.Option) (*servicediscovery.RegisterInstanceOutput, error)
	RegisterInstanceRequest(*servicediscovery.RegisterInstanceInput) (*request.Request, *servicediscovery.RegisterInstanceOutput)

	UpdateInstanceCustomHealthStatus(*servicediscovery.UpdateInstanceCustomHealthStatusInput) (*servicediscovery.UpdateInstanceCustomHealthStatusOutput, error)
	UpdateInstanceCustomHealthStatusWithContext(aws.Context, *servicediscovery.UpdateInstanceCustomHealthStatusInput, ...request.Option) (*servicediscovery.UpdateInstanceCustomHealthStatusOutput, error)
	UpdateInstanceCustomHealthStatusRequest(*servicediscovery.UpdateInstanceCustomHealthStatusInput) (*request.Request, *servicediscovery.UpdateInstanceCustomHealthStatusOutput)

	UpdateService(*servicediscovery.UpdateServiceInput) (*servicediscovery.UpdateServiceOutput, error)
	UpdateServiceWithContext(aws.Context, *servicediscovery.UpdateServiceInput, ...request.Option) (*servicediscovery.UpdateServiceOutput, error)
	UpdateServiceRequest(*servicediscovery.UpdateServiceInput) (*request.Request, *servicediscovery.UpdateServiceOutput)
}

var _ ServiceDiscoveryAPI = (*servicediscovery.ServiceDiscovery)(nil)
