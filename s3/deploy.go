package s3

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/razzkumar/PR-Automation/logger"
	"github.com/razzkumar/PR-Automation/utils"
)

// Deploy to S3 bucket
func Deploy(repo utils.ProjectInfo, sess *session.Session) error {

	dir := "./" + repo.DistFolder

	svc := s3.New(sess)

	err := CreateBucket(repo.Bucket, svc)

	if err != nil {
		logger.FailOnError(err, "Error while creating S3 bucket")
	}

	if repo.IsBuild {

		//Running build
		build()

	}

	uploader := s3manager.NewUploader(sess)

	fileList := []string{}

	filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			fileList = append(fileList, path)

			return nil
		})

	// Loop through every file and uplaod to s3
	for _, file := range fileList {
		f, _ := os.Open(file)

		key := strings.TrimPrefix(file, dir)
		key = strings.Replace(key, repo.DistFolder, "", -1)
		fileContentType := utils.GetFileType(file)

		_, err := uploader.Upload(&s3manager.UploadInput{
			Bucket:      aws.String(repo.Bucket),
			Key:         aws.String(key),
			ContentType: aws.String(fileContentType),
			Body:        f,
		})

		if err != nil {
			return err
		}
		fmt.Println("Uploading... " + key)
	}

	fmt.Println("\n\n" + strconv.Itoa(len(fileList)) + " Files Uploaded Successfully. 🎉 🎉 🎉")

	// cloudfrontDistributionID := os.Getenv("CLOUDFRONT_ID")

	// if cloudfrontDistributionID==""

	url := utils.GetURL(repo.Bucket)
	fmt.Println("URL : ", url)

	fmt.Println("removing dist files")
	os.RemoveAll(dir)

	return nil
}

func build() {

	buildCmd := os.Getenv("BUILD_COMMAND")

	if strings.Contains(buildCmd, "npm") {

		fmt.Println("npm install ....")
		utils.RunCommand("npm install")

		fmt.Printf("------------ %s ----------", buildCmd)
		utils.RunCommand(buildCmd)

	} else {

		fmt.Println("yarn install ....")
		utils.RunCommand("yarn")

		if buildCmd == "" {
			fmt.Println("----------- yarn build ----------")
			utils.RunCommand("yarn build")
		} else {
			fmt.Printf("----------- %s -----------", buildCmd)
			utils.RunCommand(buildCmd)
		}
	}
}
