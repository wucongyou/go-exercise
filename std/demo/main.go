package main

import (
	"github.com/yeeuu/echoic"
)

func main() {
	e := echoic.New()
	e.SetDebug(true)
	e.Run("127.0.0.1:4321")
}
