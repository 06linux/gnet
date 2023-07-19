package main

import (
	"fmt"
	"gnet/gbase"
)

func main() {
	fmt.Println("test db")

	gbase.Redis.Test()
	gbase.Mongo.Test()
}
