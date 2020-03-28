package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/razzkumar/PR-Automation/logger"
	"github.com/razzkumar/PR-Automation/s3"
)

func main() {
	var action string

	flag.StringVar(&action, "action", "create", "It's create or delete s3 bucket")

	flag.Parse()

	prBranch := os.Getenv("PR_BRANCH")

	if prBranch == "" {
		logger.FailOnNoFlag("PR_BRANCH not set")
	}

	prNum := os.Getenv("PR_NUMBER")

	if prNum == "" {
		logger.FailOnNoFlag("PR_NUMBER not set")
	}

	bucket := strings.ToLower(prBranch + ".PR" + prNum + ".autoDeploy")
	fmt.Println("bucket", bucket)
	// Getting session of aws

	sess := s3.GetSession()

	switch action {
	case "create":
		err := s3.Deploy(bucket, sess)
		logger.FailOnError(err, "Error on Deployment")
	case "delete":
		err := s3.Delete(bucket, sess)
		logger.FailOnError(err, "Error while Delete")
	default:
		err := fmt.Errorf("Nothing to do")
		logger.FailOnError(err, "Default case")
	}
}
