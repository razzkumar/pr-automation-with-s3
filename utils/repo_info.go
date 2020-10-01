package utils

import (
	"os"
	"strconv"
	"strings"

	"github.com/razzkumar/PR-Automation/logger"
)

type ProjectInfo struct {
	PrNumber   int
	RepoOwner  string
	Branch     string
	RepoName   string
	DistFolder string
	Bucket     string
	IsBuild    bool
}

func getDistFolder() string {

	assestFolder := os.Getenv("SRC_FOLDER")
	//setting build folder default
	if assestFolder == "" {
		assestFolder = "build"
	}

	distDir := strings.Replace(assestFolder, "./", "", -1)

	return distDir
}

func GetPRInfo(repo ProjectInfo) ProjectInfo {

	prEvent := GetPREvent()
	repo.Branch = prEvent.PullRequest.Head.GetRef()

	repo.PrNumber = prEvent.GetNumber()

	prNumInit := strconv.Itoa(repo.PrNumber)

	bucket := strings.ToLower(repo.Branch + ".PR" + prNumInit + ".auto-deploy")

	repo.Bucket = bucket

	repo.RepoOwner = prEvent.Repo.Owner.GetLogin()
	repo.RepoName = prEvent.Repo.GetName()

	repo.DistFolder = getDistFolder()

	isBuild := os.Getenv("IS_BUILD")

	if isBuild == "" || isBuild == "true" {
		repo.IsBuild = true
	}

	return repo
}

func GetInfo(repo ProjectInfo, action string) ProjectInfo {

	// setting bucket
	bucket := os.Getenv("AWS_S3_BUCKET")
	if action == "deploy" && bucket == "" {
		logger.FailOnNoFlag("AWS_S3_BUCKET is not set:")
	}
	repo.Bucket = strings.ToLower(bucket + ".auto-deploy")

	// setting dist folder
	repo.DistFolder = getDistFolder()

	isBuild := os.Getenv("IS_BUILD")

	if isBuild == "" || isBuild == "true" {
		repo.IsBuild = true
	}

	return repo
}
