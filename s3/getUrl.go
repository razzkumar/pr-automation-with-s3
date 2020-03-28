package s3

import "os"

func GetURL(bucket string) string {
	region := os.Getenv("AWS_REGION")
	url := "http://" + bucket + ".s3-website." + region + ".amazonaws.com/"
	return url
}
