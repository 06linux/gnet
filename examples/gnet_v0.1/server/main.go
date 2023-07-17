package main

import (
	"gnet/gbase"
	"gnet/gcore"
)

func main() {

	server := gcore.NewServer()
	server.Start()

	gbase.Signal.Waiting()
	// gbase.Time.GetTime()
}
