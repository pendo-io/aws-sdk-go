// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build go1.10,integration

package snowball_test

import (
	"context"
	"testing"
	"time"

	"github.com/pendo-io/aws-sdk-go/aws"
	"github.com/pendo-io/aws-sdk-go/aws/awserr"
	"github.com/pendo-io/aws-sdk-go/aws/request"
	"github.com/pendo-io/aws-sdk-go/awstesting/integration"
	"github.com/pendo-io/aws-sdk-go/service/snowball"
)

var _ aws.Config
var _ awserr.Error
var _ request.Request

func TestInteg_00_DescribeAddresses(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-west-2")
	svc := snowball.New(sess)
	params := &snowball.DescribeAddressesInput{}
	_, err := svc.DescribeAddressesWithContext(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
