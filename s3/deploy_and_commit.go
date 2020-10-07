package s3

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/razzkumar/PR-Automation/github"
	"github.com/razzkumar/PR-Automation/logger"
	"github.com/razzkumar/PR-Automation/utils"
)

// Deploy to S3 bucket
func DeployAndComment(repo utils.ProjectInfo, sess *session.Session) error {
	// TODO notity to slack
	err := Deploy(repo, sess)

	if err != nil {
		logger.FailOnError(err, "Error While Deploying to s3")
	}

	url := utils.GetURL(repo.Bucket)

	err = gh.Comment(url, repo)

	if err != nil {
		return err
	}
	return nil
}
