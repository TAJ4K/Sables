package main

import (
	"encoding/json"
	"fmt"

	_ "embed"
	"github.com/joho/godotenv"

	"github.com/tidwall/gjson"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

//go:embed .env
var env string

func rekognize(image []byte) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(getDotEnvVar("AWS_ACCESS_KEY_ID"), getDotEnvVar("AWS_SECRET_ACCESS_KEY"), ""),
	}))

	svc := rekognition.New(sess)

	input := &rekognition.DetectLabelsInput{
		Image: &rekognition.Image{
			Bytes: image,
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
		sendToast("Detected object: " + highestConfObject)
	}
}

func getDotEnvVar(key string) string {
	damap, err := godotenv.Unmarshal(env)
	if err != nil {
		panic(err)
	}

	return damap[key]
}
