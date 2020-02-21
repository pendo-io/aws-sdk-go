// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package piiface provides an interface to enable mocking the AWS Performance Insights service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package piiface

import (
	"github.com/pendo-io/aws-sdk-go/aws"
	"github.com/pendo-io/aws-sdk-go/aws/request"
	"github.com/pendo-io/aws-sdk-go/service/pi"
)

// PIAPI provides an interface to enable mocking the
// pi.PI service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Performance Insights.
//    func myFunc(svc piiface.PIAPI) bool {
//        // Make svc.DescribeDimensionKeys request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := pi.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockPIClient struct {
//        piiface.PIAPI
//    }
//    func (m *mockPIClient) DescribeDimensionKeys(input *pi.DescribeDimensionKeysInput) (*pi.DescribeDimensionKeysOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockPIClient{}
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
type PIAPI interface {
	DescribeDimensionKeys(*pi.DescribeDimensionKeysInput) (*pi.DescribeDimensionKeysOutput, error)
	DescribeDimensionKeysWithContext(aws.Context, *pi.DescribeDimensionKeysInput, ...request.Option) (*pi.DescribeDimensionKeysOutput, error)
	DescribeDimensionKeysRequest(*pi.DescribeDimensionKeysInput) (*request.Request, *pi.DescribeDimensionKeysOutput)

	GetResourceMetrics(*pi.GetResourceMetricsInput) (*pi.GetResourceMetricsOutput, error)
	GetResourceMetricsWithContext(aws.Context, *pi.GetResourceMetricsInput, ...request.Option) (*pi.GetResourceMetricsOutput, error)
	GetResourceMetricsRequest(*pi.GetResourceMetricsInput) (*request.Request, *pi.GetResourceMetricsOutput)
}

var _ PIAPI = (*pi.PI)(nil)
