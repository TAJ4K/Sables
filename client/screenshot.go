package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/moutend/go-hook/pkg/types"

	"github.com/kbinani/screenshot"
)

func screenshotThis(e1 types.MouseEvent, e2 types.MouseEvent) {
	//get the mouse location
	x1, y1 := getMouseLoc(e1)
	x2, y2 := getMouseLoc(e2)

	fmt.Println(x1, y1, x2, y2)
	//get the screenshot
	img, err := screenshot.Capture(x1, y1, x2-x1, y2-y1)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(getAppData() + "\\screenshot.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)

	rekognize()
}

func getMouseLoc(m types.MouseEvent) (int, int) {
	return int(m.X), int(m.Y)
}
