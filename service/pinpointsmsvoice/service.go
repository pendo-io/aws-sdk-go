// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointsmsvoice

import (
	"github.com/pendo-io/aws-sdk-go/aws"
	"github.com/pendo-io/aws-sdk-go/aws/client"
	"github.com/pendo-io/aws-sdk-go/aws/client/metadata"
	"github.com/pendo-io/aws-sdk-go/aws/request"
	"github.com/pendo-io/aws-sdk-go/aws/signer/v4"
	"github.com/pendo-io/aws-sdk-go/private/protocol/restjson"
)

// PinpointSMSVoice provides the API operation methods for making requests to
// Amazon Pinpoint SMS and Voice Service. See this package's package overview docs
// for details on the service.
//
// PinpointSMSVoice methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type PinpointSMSVoice struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "Pinpoint SMS Voice" // Name of service.
	EndpointsID = "sms-voice.pinpoint" // ID to lookup a service endpoint with.
	ServiceID   = "Pinpoint SMS Voice" // ServiceID is a unique identifer of a specific service.
)

// New creates a new instance of the PinpointSMSVoice client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a PinpointSMSVoice client from just a session.
//     svc := pinpointsmsvoice.New(mySession)
//
//     // Create a PinpointSMSVoice client with additional configuration
//     svc := pinpointsmsvoice.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *PinpointSMSVoice {
	c := p.ClientConfig(EndpointsID, cfgs...)
	if c.SigningNameDerived || len(c.SigningName) == 0 {
		c.SigningName = "sms-voice"
	}
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *PinpointSMSVoice {
	svc := &PinpointSMSVoice{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2018-09-05",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restjson.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a PinpointSMSVoice operation and runs any
// custom request initialization.
func (c *PinpointSMSVoice) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
