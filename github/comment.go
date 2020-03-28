package gh

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/google/go-github/v30/github"
	"github.com/razzkumar/PR-Automation/logger"
	"golang.org/x/oauth2"
)

func Comment(url string) {

	repo := os.Getenv("GH_REPO")

	if repo == "" {

		logger.FailOnNoFlag("Unbale to load repo name ")
	}
	owner := os.Getenv("REPO_USER")

	comment := "Visit: " + url

	num, err := strconv.Atoi(os.Getenv("PR_NUMBER"))
	if err != nil {
		logger.FailOnError(err, "Error While Parsing PR number")
	}

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GH_ACCSS_TOKEN")},
	)

	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	pullRequestReviewRequest := &github.PullRequestReviewRequest{Body: &comment, Event: github.String("COMMENT")}

	//client.PullRequests.CreateComment(ctx, owner, repo, num, pullRequestReviewRequest)
	pullRequestReview, _, err := client.PullRequests.CreateReview(ctx, owner, repo, num, pullRequestReviewRequest)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("github-Commit: Created GitHub PR Review comment", pullRequestReview.ID)
}
