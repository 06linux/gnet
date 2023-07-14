package main

import (
	"gnet/gcore"
	"gnet/gutil"
)

func main() {

	server := gcore.NewServer()
	server.Start()

	gutil.Signal.Waiting()
	// gutil.Time.GetTime()
}
