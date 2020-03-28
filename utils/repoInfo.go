package utils

import (
	"os"
	"strconv"
	"strings"

	"github.com/razzkumar/PR-Automation/logger"
)

type GithubInfo struct {
	PrNumber  int
	RepoOwner string
	Branch    string
	RepoName  string
}

func GetPRInfo() GithubInfo {

	repo := GithubInfo{}

	// It's on the form of "refs/pull/1/merge"
	_ghref := os.Getenv("GITHUB_REF")

	if _ghref == "" {
		logger.FailOnNoFlag("Unable to load GITHUB_REF")
	}

	ghref := strings.Split(_ghref, "/")
	prNum, err := strconv.Atoi(ghref[2])

	if err != nil {
		logger.FailOnError(err, "Error While Parsing PR number")
	}

	repo.PrNumber = prNum

	prBranch := os.Getenv("GITHUB_HEAD_REF")

	if prBranch == "" {
		logger.FailOnNoFlag("PR_BRANCH not set")
	}
	repo.Branch = prBranch

	// It's on the form of "razzkumar/ftodo"

	_ghRepo := os.Getenv("GITHUB_REPOSITORY")

	if _ghRepo == "" {
		logger.FailOnNoFlag("Unable to parse GITHUB_REPOSITORY")
	}

	ghRepo := strings.Split(_ghRepo, "/")

	repo.RepoOwner = ghRepo[0]
	repo.RepoName = ghRepo[1]

	return repo
}
