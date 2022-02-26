package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/joho/godotenv"

	"github.com/tidwall/gjson"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

func rekognize() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(getDotEnvVar("AWS_ACCESS_KEY_ID"), getDotEnvVar("AWS_SECRET_ACCESS_KEY"), ""),
	}))

	image, _ := os.Open(getAppData() + "\\screenshot.png")
	imageBytes, _ := io.ReadAll(image)

	svc := rekognition.New(sess)

	input := &rekognition.DetectLabelsInput{
		Image: &rekognition.Image{
			Bytes: imageBytes,
		},
	}

	result, err := svc.DetectLabels(input)
	if err != nil {
		panic(err)
	}

	output, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	highestConfObject := gjson.Get(string(output), "Labels.0.Name").String()
	fmt.Println(highestConfObject)
	if highestConfObject != "" {
		tts(highestConfObject)
	}
}

func getDotEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	
	return os.Getenv(key)
}
