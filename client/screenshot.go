package main

import (
	"fmt"
	"image/png"
	"bytes"

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
		return
	}

	buff := new(bytes.Buffer)
	
	err = png.Encode(buff, img)
	if err != nil {
		panic(err)
	}

	rekognize(buff.Bytes())
}

func getMouseLoc(m types.MouseEvent) (int, int) {
	return int(m.X), int(m.Y)
}
