package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	bucketName = "golang-image-api-bucket-1"
	regionName = "us-east-1"
)

var (
	grabbedLink string
	s3session   *s3.S3
)

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(regionName)},
	)

	if err != nil {
		os.Exit(1)
	}

	s3session = s3.New(sess)
}

func convertImgToBytes(url string) (bytes []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return bytes, err
	}

	defer resp.Body.Close()

	bytes, err = io.ReadAll(resp.Body)
	return bytes, err
}

func handler() (events.APIGatewayProxyResponse, error) {
	img, _ := convertImgToBytes(grabbedLink)

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "image/jpeg",
		},
		Body:            base64.StdEncoding.EncodeToString(img),
		IsBase64Encoded: true,
	}

	return response, nil
}

func listObjects() (obj *s3.ListObjectsV2Output) {
	resp, err := s3session.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})

	if err != nil {
		log.Println("err in the bucket listing")
		os.Exit(1)
	}

	for _, item := range resp.Contents {
		grabbedLink = fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, regionName, *item.Key)
	}

	return resp
}

func main() {
	listObjects()
	lambda.Start(handler)
}
