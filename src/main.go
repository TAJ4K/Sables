package main

import(
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/key"

	"github.com/kbinani/screenshot"
)

func main(){
	keyAltPressed, keyHPressed := false, false
	var e1, e2 mouse.Event = mouse.Event{}, mouse.Event{}
	app.Main(func(a app.App){ 
		for e := range a.Events(){
			switch e := a.Filter(e).(type) {
				case key.Code:
					switch e {
						case key.CodeLeftAlt:
							//start the key combo if alt pressed
							keyAltPressed = true
						case key.CodeH:
							//marks h as pressed
							keyHPressed = true
						default:
							//exit the combo if any other key is pressed
							keyAltPressed = false
							keyHPressed = false
					}
				case mouse.Event:
					if keyAltPressed && keyHPressed && (e2 == mouse.Event{}) {
						// if the keys have been pressed, and e2 isnt assigned yet
						e1 = e
					} else if keyAltPressed && keyHPressed && (e1 == mouse.Event{}) {
						// if the keys have been pressed, and e1 is assigned
						e2 = e
						screenshotThis(e1, e2)
					}
			}
		}
	}) 
}

func getMouseLoc(e mouse.Event) (int, int){
	return int(e.X), int(e.Y)
}

func screenshotThis(e1 mouse.Event, e2 mouse.Event) {
	//get the mouse location
	x1, y1 := getMouseLoc(e1)
	x2, y2 := getMouseLoc(e2)
	//get the screenshot
	img, err := screenshot.Capture(x1, y1, x2-x1, y2-y1)
	if err != nil {
		panic(err)
	}
}