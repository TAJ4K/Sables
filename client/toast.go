package main

import (
	"io"
	"net/http"
	"os"

	"github.com/go-toast/toast"
)

var basePathG string

func basePath() {
	basePath, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	basePathG = basePath
}

func initAppData(){
	os.MkdirAll(basePathG+"/Sables", 0777)

	out, err := os.Create(basePathG+"/Sables/icon.png")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	//discord CDN
	resp, err := http.Get("https://cdn.discordapp.com/attachments/819783862762209299/947329323558395924/sablesico.png")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}

func sendToast(message string) {
	notification := toast.Notification{
		AppID:   "Sables",
		Title:   "\u200b",
		Message: message,
		Icon:    basePathG+"/Sables/icon.png",
	}

	notification.Push()
}