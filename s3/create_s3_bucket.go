package s3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetStaticSiteStatus(bucket string, svc *s3.S3) (*s3.GetBucketWebsiteOutput, error) {
	status, err := svc.GetBucketWebsite(&s3.GetBucketWebsiteInput{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		// Check for the NoSuchWebsiteConfiguration error code telling us
		// that the bucket does not have a website configured.
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "NoSuchWebsiteConfiguration" {
			EnableStaticHosting(bucket, svc)
		} else {
			return nil, fmt.Errorf("Unable to get bucket website config, %v", err)
		}
	}
	fmt.Println("Bucket Status", bucket)
	return status, nil
}

func CreateBucket(bucket string, svc *s3.S3) error {
	_, err := svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "BucketAlreadyOwnedByYou" {
			fmt.Printf("bucket %s is Already Exist\n", bucket)
		} else {
			return err
		}
	}
	GetStaticSiteStatus(bucket, svc)
	isPublic, _ := GetPolicy(bucket, svc)
	if !isPublic {
		fmt.Println("Attaching policy")
		AttachPolicy(bucket, svc)
	}
	return nil
}
