package s3control

import (
	"github.com/pendo-io/aws-sdk-go/aws/client"
	"github.com/pendo-io/aws-sdk-go/internal/s3err"
)

func init() {
	initClient = defaultInitClientFn
}

func defaultInitClientFn(c *client.Client) {
	c.Handlers.UnmarshalError.PushBackNamed(s3err.RequestFailureWrapperHandler())
}
