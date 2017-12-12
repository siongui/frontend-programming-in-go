package main

import (
	. "github.com/siongui/godom"
)

const (
	LEFT  = 37
	UP    = 38
	RIGHT = 39
	DOWN  = 40
)

func main() {
	info := Document.QuerySelector("#info")

	Window.AddEventListener("keyup", func(event Event) {
		keycode := event.KeyCode()
		if keycode == LEFT {
			info.Set("innerHTML", "left arrow key")
		}
		if keycode == UP {
			info.Set("innerHTML", "up arrow key")
		}
		if keycode == RIGHT {
			info.Set("innerHTML", "right arrow key")
		}
		if keycode == DOWN {
			info.Set("innerHTML", "down arrow key")
		}
	}, false)
}
