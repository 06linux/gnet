package main

import (
	"gnet/gbase"
	"gnet/gcore"
)

func main() {

	client := gcore.NewClient()
	client.Connect("127.0.0.1", 8000)

	gbase.Signal.Waiting()
}
