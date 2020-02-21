// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elastictranscoder

import (
	"github.com/pendo-io/aws-sdk-go/aws"
	"github.com/pendo-io/aws-sdk-go/aws/client"
	"github.com/pendo-io/aws-sdk-go/aws/client/metadata"
	"github.com/pendo-io/aws-sdk-go/aws/request"
	"github.com/pendo-io/aws-sdk-go/aws/signer/v4"
	"github.com/pendo-io/aws-sdk-go/private/protocol/restjson"
)

// ElasticTranscoder provides the API operation methods for making requests to
// Amazon Elastic Transcoder. See this package's package overview docs
// for details on the service.
//
// ElasticTranscoder methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type ElasticTranscoder struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "elastictranscoder"  // Name of service.
	EndpointsID = ServiceName          // ID to lookup a service endpoint with.
	ServiceID   = "Elastic Transcoder" // ServiceID is a unique identifer of a specific service.
)

// New creates a new instance of the ElasticTranscoder client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a ElasticTranscoder client from just a session.
//     svc := elastictranscoder.New(mySession)
//
//     // Create a ElasticTranscoder client with additional configuration
//     svc := elastictranscoder.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *ElasticTranscoder {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *ElasticTranscoder {
	svc := &ElasticTranscoder{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2012-09-25",
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

// newRequest creates a new request for a ElasticTranscoder operation and runs any
// custom request initialization.
func (c *ElasticTranscoder) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
