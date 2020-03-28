package main

import (
	"flag"
	"fmt"

	"github.com/razzkumar/PR-Automation/logger"
	"github.com/razzkumar/PR-Automation/s3"
)

func main() {
	var action, bucket string

	flag.StringVar(&action, "action", "create", "It's create or delete s3 bucket")
	flag.StringVar(&bucket, "bucket", "", "Name of the s3 bucket to create")

	flag.Parse()

	if bucket == "" {
		logger.FailOnNoFlag("s3 bucket name required")
	}

	// Getting session of aws

	sess := s3.GetSession()

	data := s3.Data{
		DistDir:    "./build",
		BucketName: bucket,
	}

	switch action {
	case "create":
		err := s3.Deploy(data, sess)
		logger.FailOnError(err, "Error on Deployment")
	case "delete":
		err := s3.Delete(data.BucketName, sess)
		logger.FailOnError(err, "Error while Delete")
	default:
		err := fmt.Errorf("Nothing to do")
		logger.FailOnError(err, "Default case")
	}
}
