package main

import (
	"io"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func tts(){
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(getDotEnvVar("AWS_ACCESS_KEY_ID"), getDotEnvVar("AWS_SECRET_ACCESS_KEY"), ""),
	}))

	svc := polly.New(sess)

	input := &polly.SynthesizeSpeechInput{OutputFormat: aws.String("mp3"), Text: aws.String("Hello World"), VoiceId: aws.String("Matthew")}

	result, err := svc.SynthesizeSpeech(input)
	if err != nil {
		panic(err)
	}

	mp3File := "speech" + ".mp3"

	outFile, err := os.Create(getAppData() + "\\" + mp3File)
	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, result.AudioStream)
	if err != nil {
		panic(err)
	}
}

func playMP3() {
	file, err := os.Open(getAppData() + "\\speech.mp3")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	streamer, format, err := mp3.Decode(file)
	if err != nil {
		panic(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}