package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/go-github/v30/github"
	"github.com/razzkumar/PR-Automation/logger"
)

func ParseGithubEvent() interface{} {

	event := os.Getenv("GITHUB_EVENT_NAME")
	filePath := os.Getenv("GITHUB_EVENT_PATH")
	// Open our jsonFile
	jsonFile, err := os.Open(filePath)

	// if we os.Open returns an error then handle it
	if err != nil {
		logger.FailOnError(err, "Fail to read event file")
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	parsedData, err := github.ParseWebHook(event, byteValue)

	if err != nil {
		log.Fatal(err)
	}

	return parsedData

}

func GetPREvent() *github.PullRequestEvent {
	parsedData := ParseGithubEvent()
	return parsedData.(*github.PullRequestEvent)
}

func GetPushEvent() *github.PushEvent {
	parsedData := ParseGithubEvent()
	return parsedData.(*github.PushEvent)
}
