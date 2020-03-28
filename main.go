package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/razzkumar/PR-Automation/logger"
	"github.com/razzkumar/PR-Automation/s3"
	"github.com/razzkumar/PR-Automation/utils"
)

func main() {
	var action string

	// Getting action wether delete or create
	flag.StringVar(&action, "action", "create", "It's create or delete s3 bucket")
	flag.Parse()

	repo := utils.GetPRInfo()
	fmt.Println((repo))
	fmt.Println(repo.PrNumber)
	prNum := strconv.Itoa(repo.PrNumber)
	fmt.Printf("%v : %T", prNum, prNum)
	bucket := strings.ToLower(repo.Branch + ".PR" + prNum + ".auto-deploy")
	fmt.Println("bucket", bucket)
	//// Getting session of aws

	sess := s3.GetSession()

	switch action {
	case "create":
		err := s3.Deploy(bucket, repo, sess)
		logger.FailOnError(err, "Error on Deployment")
	case "delete":
		err := s3.Delete(bucket, sess)
		logger.FailOnError(err, "Error while Delete")
	default:
		err := fmt.Errorf("Nothing to do")
		logger.FailOnError(err, "Default case")
	}
}
