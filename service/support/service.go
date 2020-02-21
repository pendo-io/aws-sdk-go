// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package support

import (
	"github.com/pendo-io/aws-sdk-go/aws"
	"github.com/pendo-io/aws-sdk-go/aws/client"
	"github.com/pendo-io/aws-sdk-go/aws/client/metadata"
	"github.com/pendo-io/aws-sdk-go/aws/request"
	"github.com/pendo-io/aws-sdk-go/aws/signer/v4"
	"github.com/pendo-io/aws-sdk-go/private/protocol/jsonrpc"
)

// Support provides the API operation methods for making requests to
// AWS Support. See this package's package overview docs
// for details on the service.
//
// Support methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type Support struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "support"   // Name of service.
	EndpointsID = ServiceName // ID to lookup a service endpoint with.
	ServiceID   = "Support"   // ServiceID is a unique identifer of a specific service.
)

// New creates a new instance of the Support client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a Support client from just a session.
//     svc := support.New(mySession)
//
//     // Create a Support client with additional configuration
//     svc := support.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *Support {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *Support {
	svc := &Support{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2013-04-15",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWSSupport_20130415",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a Support operation and runs any
// custom request initialization.
func (c *Support) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
