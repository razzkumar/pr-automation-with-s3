package cloudfront

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
)

func Invalidation(did string, sess *session.Session) error {

	if did == "" {
		return nil
	}

	client := cloudfront.New(sess)

	now := time.Now()

	invalidaitonInput := cloudfront.CreateInvalidationInput{
		DistributionId: aws.String(did),
		InvalidationBatch: &cloudfront.InvalidationBatch{
			CallerReference: aws.String(
				fmt.Sprintf("goinvali%s", now.Format("2006/01/02,15:04:05"))),
			Paths: &cloudfront.Paths{
				Quantity: aws.Int64(1),
				Items: []*string{
					aws.String("/*"),
				},
			},
		},
	}

	resp, err := client.CreateInvalidation(&invalidaitonInput)

	if err != nil {
		return fmt.Errorf("Invalidation failed. err:%v", err.Error())
	}

	fmt.Printf("Invalidation:%v\n", resp)

	return nil

}
