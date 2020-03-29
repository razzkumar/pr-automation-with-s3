package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/razzkumar/PR-Automation/logger"
	"github.com/razzkumar/PR-Automation/s3"
	"github.com/razzkumar/PR-Automation/utils"
)

func main() {

	// Setting env variable
	awsRegion := os.Getenv("AWS_REGION")

	if awsRegion == "" {
		err := os.Setenv("AWS_REGION", "us-east-2")
		if err != nil {
			logger.FailOnError(err, "Fail to set AWS_REGION")
		}
	}

	var action string
	var repo utils.ProjectInfo
	// Getting action wether delete or create
	flag.StringVar(&action, "action", "", "It's create or delete s3 bucket")
	flag.Parse()

	if action == "" {
		logger.FailOnNoFlag("Please provide action what to do [deploy,delete,create]")
	}

	if os.Getenv("GITHUB_EVENT_NAME") == "pull_request" && (action == "create" || action == "delete") {
		repo = utils.GetPRInfo(repo)
	} else {
		repo = utils.GetInfo(repo, action)
	}

	// Getting session of aws
	sess := s3.GetSession()

	switch action {
	case "deploy":
		err := s3.Deploy(repo, sess)
		logger.FailOnError(err, "Error on Deployment")
	case "create":
		err := s3.DeployAndComment(repo, sess)
		logger.FailOnError(err, "Error on Deployment and commit")
	case "delete":
		err := s3.Delete(repo.Bucket, sess)
		logger.FailOnError(err, "Error while Delete")
	default:
		err := fmt.Errorf("Nothing to do")
		logger.FailOnError(err, "Default case")
	}
}
