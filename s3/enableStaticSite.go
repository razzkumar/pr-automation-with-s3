package s3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/razzkumar/PR-Automation/logger"
)

const (
	indexPage = "index.html"
	errorPage = "index.html" // If error accure redirect to index.html
)

func EnableStaticHosting(bucket string, svc *s3.S3) {
	params := s3.PutBucketWebsiteInput{
		Bucket: aws.String(bucket),
		WebsiteConfiguration: &s3.WebsiteConfiguration{
			IndexDocument: &s3.IndexDocument{
				Suffix: aws.String(indexPage),
			},
		},
	}

	// Add the error page if set on CLI
	if len(errorPage) > 0 {
		params.WebsiteConfiguration.ErrorDocument = &s3.ErrorDocument{
			Key: aws.String(errorPage),
		}
	}

	_, err := svc.PutBucketWebsite(&params)
	if err != nil {
		logger.Info(fmt.Sprintf("Unable to set bucket %q website configuration, %v", bucket, err))
	}

	fmt.Printf("Successfully set bucket %q website configuration\n", bucket)
}
