package main

import (
	"fmt"
	"os"
	"time"

	"github.com/moutend/go-hook/pkg/keyboard"
	"github.com/moutend/go-hook/pkg/mouse"
	"github.com/moutend/go-hook/pkg/types"
)

func main() {
	// _, err := os.Stat(getAppData())
	// if err != nil {
	// 	fmt.Println("Creating directory")
	// 	err = os.Mkdir(getAppData(), 0777)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// if err := keyListener(); err != nil {
	// 	panic(err)
	// }
	tts()
}

func keyListener() error {
	key1Pressed := false
	key1, key2 := "VK_LMENU", "VK_H"

	keyboardChan := make(chan types.KeyboardEvent, 100)
	if err := keyboard.Install(nil, keyboardChan); err != nil {
		return err
	}

	defer keyboard.Uninstall()

	fmt.Println("start capturing keyboard input")

	for e := range keyboardChan {
		switch e.VKCode.String() {
		case key1:
			fmt.Println(key1 + " pressed")
			key1Pressed = true
		case key2:
			fmt.Println(key2 + " pressed")
			if key1Pressed {
				if err := mouseListener(); err != nil {
					return err
				}
				key1Pressed = false
			}
		default:
			//resets state if key combo is blocked
			key1Pressed = false
		}
	}
	return nil
}

func mouseListener() error {
	var mousePos1 types.MouseEvent

	mouseChan := make(chan types.MouseEvent)

	if err := mouse.Install(nil, mouseChan); err != nil {
		return err
	}

	defer mouse.Uninstall()

	for {
		select {
		case <-time.After(10 * time.Second):
			fmt.Println("Timeout")
			return nil
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

func getAppData() string {
	basePath, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	return basePath + "\\Sables"
}
