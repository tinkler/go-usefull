package main

import (
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	for {
		// 等待5秒
		time.Sleep(5 * time.Second)

		robotgo.MouseSleep = 100

		robotgo.ScrollDir(10, "up")
		robotgo.ScrollDir(20, "right")

		robotgo.Scroll(0, -10)
		robotgo.Scroll(100, 0)

		robotgo.MilliSleep(100)
		robotgo.ScrollSmooth(-10, 6)
		// robotgo.ScrollRelative(10, -100)

		robotgo.Move(10, 20)
		robotgo.MoveRelative(0, -10)
		robotgo.DragSmooth(10, 10)

		robotgo.Click("wheelRight")
		robotgo.Click("left", true)
		robotgo.MoveSmooth(100, 200, 1.0, 10.0)

		robotgo.Toggle("left")
		robotgo.Toggle("left", "up")
	}
}
