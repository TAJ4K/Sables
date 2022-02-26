package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/moutend/go-hook/pkg/keyboard"
	"github.com/moutend/go-hook/pkg/mouse"
	"github.com/moutend/go-hook/pkg/types"

	"github.com/kbinani/screenshot"
)

func main(){
	_, err := os.Stat(getAppData()) 
	if err != nil {
		fmt.Println("Creating directory")
		err = os.Mkdir(getAppData(), 0777)
		if err != nil {
			panic(err)
		}
	}

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error{
	keyAltPressed := false
	keyAlt, keyH := "VK_LMENU", "VK_H"

	keyboardChan := make(chan types.KeyboardEvent, 100)
	if err := keyboard.Install(nil, keyboardChan); err != nil {
		return err
	}

	defer keyboard.Uninstall()

	fmt.Println("start capturing keyboard input")

	for {
		select {
		case e := <-keyboardChan:
			if e.VKCode.String() == keyAlt {
				fmt.Println("Alt pressed")
				keyAltPressed = true
			}
			if e.VKCode.String() == keyH && keyAltPressed {
				fmt.Println("H pressed")
				if err := mouseListener(); err != nil {
					return err
				}
				keyAltPressed = false
			}
		}
	}
}

func mouseListener() error{
	var mousePos1 types.MouseEvent

	mouseChan := make(chan types.MouseEvent)

	if err := mouse.Install(nil, mouseChan); err != nil {
		return err
	}

	defer mouse.Uninstall()

	for {
		select {
			case e := <-mouseChan:
				if e.Message.String() == "Message(513)" {
					fmt.Println("Left button pressed")
					mousePos1 = e
				} else if e.Message.String() == "Message(514)" {
					fmt.Println("Left button released")
					screenshotThis(mousePos1, e)
					return nil
				}
		}
	}
	
}

func getMouseLoc(m types.MouseEvent) (int, int){
	return int(m.X), int(m.Y)
}

func getAppData() string {
	basePath, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	return basePath + "\\Sable"
}

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
}