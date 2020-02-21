// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build go1.10,integration

package sns_test

import (
	"context"
	"testing"
	"time"

	"github.com/pendo-io/aws-sdk-go/aws"
	"github.com/pendo-io/aws-sdk-go/aws/awserr"
	"github.com/pendo-io/aws-sdk-go/aws/request"
	"github.com/pendo-io/aws-sdk-go/awstesting/integration"
	"github.com/pendo-io/aws-sdk-go/service/sns"
)

var _ aws.Config
var _ awserr.Error
var _ request.Request

func TestInteg_00_ListTopics(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-west-2")
	svc := sns.New(sess)
	params := &sns.ListTopicsInput{}
	_, err := svc.ListTopicsWithContext(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
func TestInteg_01_Publish(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-west-2")
	svc := sns.New(sess)
	params := &sns.PublishInput{
		Message:  aws.String("hello"),
		TopicArn: aws.String("fake_topic"),
	}
	_, err := svc.PublishWithContext(ctx, params)
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
