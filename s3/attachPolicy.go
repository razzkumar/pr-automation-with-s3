package s3

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

func AttachPolicy(bucket string, svc *s3.S3) error {

	publicRead := map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Sid":       "PublicReadGetObject",
				"Effect":    "Allow",
				"Principal": "*",
				"Action": []string{
					"s3:GetObject",
				},
				"Resource": []string{
					fmt.Sprintf("arn:aws:s3:::%s/*", bucket),
				},
			},
		},
	}

	policy, err := json.Marshal(publicRead)
	if err != nil {
		return err
	}
	_, err = svc.PutBucketPolicy(&s3.PutBucketPolicyInput{
		Bucket: aws.String(bucket),
		Policy: aws.String(string(policy)),
	})

	if err != nil {
		return err
	}

	fmt.Printf("Successfully set bucket %q's policy\n", bucket)
	return nil
}

func GetPolicy(bucket string, svc *s3.S3) (bool, error) {
	result, err := svc.GetBucketPolicy(&s3.GetBucketPolicyInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		// Special error handling for the when the bucket doesn't
		// exists so we can give a more direct error message from the CLI.
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				return false, fmt.Errorf("Bucket %q does not exist.", bucket)

			case "NoSuchBucketPolicy":
				return false, fmt.Errorf("Bucket %q does not have a policy.", bucket)
			}
		}
		return false, fmt.Errorf("Unable to get bucket %q policy, %v.", bucket, err)
	}

	isPublic := strings.Contains(result.String(), "PublicReadGetObject")
	fmt.Println("isPublic", isPublic)
	if !isPublic {
		return false, fmt.Errorf("Bucket is not Public")
	}
	return true, nil
}
