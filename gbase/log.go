package gbase

import (
	"fmt"
	"log"
	// "os"
)

type baseLog struct{}

func (baseLog) Init() {

	fmt.Println("LogInit ...")

	// file, err := os.OpenFile("./gnet.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// if err != nil {
	// 	fmt.Println("open log file failed, err:", err)
	// 	return
	// }
	// log.SetOutput(file)
	// log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetFlags(log.Llongfile)
}
