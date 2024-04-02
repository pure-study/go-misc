package main

import (
	"github.com/go-vgo/robotgo"
)

func main() {
	robotgo.MouseSleep = 100

	robotgo.Move(10, 20)
}
