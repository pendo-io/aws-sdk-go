// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build go1.10,integration

package elastictranscoder_test

import (
	"context"
	"testing"
	"time"

	"github.com/pendo-io/aws-sdk-go/aws"
	"github.com/pendo-io/aws-sdk-go/aws/awserr"
	"github.com/pendo-io/aws-sdk-go/aws/request"
	"github.com/pendo-io/aws-sdk-go/awstesting/integration"
	"github.com/pendo-io/aws-sdk-go/service/elastictranscoder"
)

var _ aws.Config
var _ awserr.Error
var _ request.Request

func TestInteg_00_ListPresets(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-west-2")
	svc := elastictranscoder.New(sess)
	params := &elastictranscoder.ListPresetsInput{}
	_, err := svc.ListPresetsWithContext(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
func TestInteg_01_ReadJob(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-west-2")
	svc := elastictranscoder.New(sess)
	params := &elastictranscoder.ReadJobInput{
		Id: aws.String("fake_job"),
	}
	_, err := svc.ReadJobWithContext(ctx, params)
	if err == nil {
		t.Fatalf("expect request to fail")
	}
	aerr, ok := err.(awserr.RequestFailure)
	if !ok {
		t.Fatalf("expect awserr, was %T", err)
	}
	if len(aerr.Code()) == 0 {
		t.Errorf("expect non-empty error code")
	}
	if v := aerr.Code(); v == request.ErrCodeSerialization {
		t.Errorf("expect API error code got serialization failure")
	}
}
