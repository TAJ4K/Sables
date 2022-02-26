package main

import (
	"io"
	"os"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

func rekognize() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(getDotEnvVar("AWS_ACCESS_KEY_ID"), getDotEnvVar("AWS_SECRET_ACCESS_KEY"), ""),
	}))

	image, _ := os.Open(getAppData() + "\\screenshot.png")
	imageBytes, _ := io.ReadAll(image)

	svc := rekognition.New(sess)

	input := &rekognition.DetectLabelsInput{
		Image: &rekognition.Image{
			Bytes: []byte(imageBytes),
		},
	}

	result, err := svc.DetectLabels(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func getDotEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}