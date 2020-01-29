// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewaymanagementapi

import (
	"github.com/pendo-io/aws-sdk-go/aws"
	"github.com/pendo-io/aws-sdk-go/aws/awsutil"
	"github.com/pendo-io/aws-sdk-go/aws/request"
	"github.com/pendo-io/aws-sdk-go/private/protocol"
	"github.com/pendo-io/aws-sdk-go/private/protocol/restjson"
)

const opPostToConnection = "PostToConnection"

// PostToConnectionRequest generates a "aws/request.Request" representing the
// client's request for the PostToConnection operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See PostToConnection for more information on using the PostToConnection
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the PostToConnectionRequest method.
//    req, resp := client.PostToConnectionRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/apigatewaymanagementapi-2018-11-29/PostToConnection
func (c *ApiGatewayManagementApi) PostToConnectionRequest(input *PostToConnectionInput) (req *request.Request, output *PostToConnectionOutput) {
	op := &request.Operation{
		Name:       opPostToConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/@connections/{connectionId}",
	}

	if input == nil {
		input = &PostToConnectionInput{}
	}

	output = &PostToConnectionOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Swap(restjson.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	return
}

// PostToConnection API operation for AmazonApiGatewayManagementApi.
//
// Sends the provided data to the specified connection.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AmazonApiGatewayManagementApi's
// API operation PostToConnection for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeGoneException "GoneException"
//   The connection with the provided id no longer exists.
//
//   * ErrCodeLimitExceededException "LimitExceededException"
//   The client is sending more than the allowed number of requests per unit of
//   time.
//
//   * ErrCodePayloadTooLargeException "PayloadTooLargeException"
//   The data has exceeded the maximum size allowed.
//
//   * ErrCodeForbiddenException "ForbiddenException"
//   The caller is not authorized to invoke this operation.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/apigatewaymanagementapi-2018-11-29/PostToConnection
func (c *ApiGatewayManagementApi) PostToConnection(input *PostToConnectionInput) (*PostToConnectionOutput, error) {
	req, out := c.PostToConnectionRequest(input)
	return out, req.Send()
}

// PostToConnectionWithContext is the same as PostToConnection with the addition of
// the ability to pass a context and additional request options.
//
// See PostToConnection for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ApiGatewayManagementApi) PostToConnectionWithContext(ctx aws.Context, input *PostToConnectionInput, opts ...request.Option) (*PostToConnectionOutput, error) {
	req, out := c.PostToConnectionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type PostToConnectionInput struct {
	_ struct{} `type:"structure" payload:"Data"`

	// ConnectionId is a required field
	ConnectionId *string `location:"uri" locationName:"connectionId" type:"string" required:"true"`

	// The data to be sent to the client specified by its connection id.
	//
	// Data is a required field
	Data []byte `type:"blob" required:"true"`
}

// String returns the string representation
func (s PostToConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PostToConnectionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PostToConnectionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PostToConnectionInput"}
	if s.ConnectionId == nil {
		invalidParams.Add(request.NewErrParamRequired("ConnectionId"))
	}
	if s.ConnectionId != nil && len(*s.ConnectionId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("ConnectionId", 1))
	}
	if s.Data == nil {
		invalidParams.Add(request.NewErrParamRequired("Data"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetConnectionId sets the ConnectionId field's value.
func (s *PostToConnectionInput) SetConnectionId(v string) *PostToConnectionInput {
	s.ConnectionId = &v
	return s
}

// SetData sets the Data field's value.
func (s *PostToConnectionInput) SetData(v []byte) *PostToConnectionInput {
	s.Data = v
	return s
}

type PostToConnectionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PostToConnectionOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PostToConnectionOutput) GoString() string {
	return s.String()
}
