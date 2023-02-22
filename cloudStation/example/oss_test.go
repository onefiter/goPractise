package example_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	ACCESS_KEY_ID     = ""
	ACCESS_KEY_SECRET = ""
	REGION            = ""
	ENDPOINT          = ""
	BUCKET_NAME       = ""
)

// BucketList接口
func TestBucketList(t *testing.T) {
	creds := credentials.NewStaticCredentials(ACCESS_KEY_ID, ACCESS_KEY_SECRET, "")
	cfg := aws.NewConfig().
		WithRegion(REGION).
		WithEndpoint(ENDPOINT).
		WithLogLevel(aws.LogDebugWithHTTPBody | aws.LogDebugWithRequestRetries).
		WithS3ForcePathStyle(true).
		WithDisableSSL(true).WithCredentials(creds)
	sess := session.Must(session.NewSession(cfg))
	svc := s3.New(sess)
	resp, err := svc.ListBuckets(&s3.ListBucketsInput{})

	if err != nil {
		panic("error")
		//fmt.Printl
	}

	for _, v := range resp.Buckets {
		fmt.Println(v)
	}
}
