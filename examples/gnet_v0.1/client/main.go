package main

import (
	"gnet/gcore"
	"gnet/gutil"
)

func main() {

	client := gcore.NewClient()
	client.Connect("127.0.0.1", 8000)

	gutil.Signal.Waiting()
}
