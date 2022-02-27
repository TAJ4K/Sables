package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gen2brain/beeep"
)

var basePathG string

func initAppData(){
	basePath, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	basePathG = basePath

	os.MkdirAll(basePath+"/Sables", 0777)

	out, err := os.Create(basePath+"/Sables/icon.png")
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
	fmt.Println(basePathG+"/Sables/icon.png")
	err := beeep.Notify("Sables", message, basePathG+"/Sables/icon.png")
	if err != nil {
		panic(err)
	}
}