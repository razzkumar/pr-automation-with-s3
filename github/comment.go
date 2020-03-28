package gh

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/v30/github"
	"github.com/razzkumar/PR-Automation/utils"
	"golang.org/x/oauth2"
)

func Comment(url string, repo utils.GithubInfo) {

	comment := "Visit: " + url

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GH_ACCSS_TOKEN")},
	)

	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	pullRequestReviewRequest := &github.PullRequestReviewRequest{Body: &comment, Event: github.String("COMMENT")}

	//client.PullRequests.CreateComment(ctx, owner, repo, num, pullRequestReviewRequest)
	pullRequestReview, _, err := client.PullRequests.CreateReview(ctx, repo.RepoOwner, repo.RepoName, repo.PrNumber, pullRequestReviewRequest)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("github-Commit: Created GitHub PR Review comment", pullRequestReview.ID)
}
