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

	assestFolder := os.Getenv("BUILD_FOLDER")

	if assestFolder == "" {
		logger.FailOnNoFlag("Unable to parse BUILD_FOLDER / source folder")
	}

	distDir := strings.Replace(assestFolder, "./", "", -1)

	return distDir
}

func GetPRInfo(repo ProjectInfo) ProjectInfo {

	//repo := ProjectInfo{}

	prBranch := os.Getenv("GITHUB_HEAD_REF")
	repo.Branch = prBranch

	// It's on the form of "refs/pull/1/merge"
	_ghref := os.Getenv("GITHUB_REF")

	if _ghref != "" {

		ghref := strings.Split(_ghref, "/")
		prNum, err := strconv.Atoi(ghref[2])

		if err != nil {
			logger.FailOnError(err, "Error While Parsing PR number")
		}

		repo.PrNumber = prNum

		prNumInit := strconv.Itoa(prNum)

		bucket := strings.ToLower(repo.Branch + ".PR" + prNumInit + ".auto-deploy")

		repo.Bucket = bucket

		//logger.FailOnNoFlag("Unable to load GITHUB_REF")
	}

	// It's on the form of "razzkumar/ftodo"

	_ghRepo := os.Getenv("GITHUB_REPOSITORY")

	if _ghRepo != "" {
		ghRepo := strings.Split(_ghRepo, "/")

		repo.RepoOwner = ghRepo[0]
		repo.RepoName = ghRepo[1]

		//logger.FailOnNoFlag("Unable to parse GITHUB_REPOSITORY")
	}

	repo.DistFolder = getDistFolder()

	isBuild := os.Getenv("IS_BUILD")

	if isBuild == "true" {
		repo.IsBuild = true
	}

	return repo
}

func GetInfo(repo ProjectInfo) ProjectInfo {

	// setting bucket
	bucket := os.Getenv("AWS_BUCKET")
	repo.Bucket = bucket

	// setting dist folder
	repo.DistFolder = getDistFolder()

	isBuild := os.Getenv("IS_BUILD")

	if isBuild == "true" {
		repo.IsBuild = true
	}

	return repo
}
