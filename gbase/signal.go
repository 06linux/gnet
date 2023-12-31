package gbase

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

type baseSignal struct{}

var Signal = &baseSignal{}

// 阻塞，直到收到 ctrl+c kill 信号则返回
func (baseSignal) Waiting() {

	// 阻塞,否则主Go退出， listenner的go将会退出
	c := make(chan os.Signal, 1)

	// 监听指定信号 ctrl+c kill信号
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c

	log.Println("[Info] Waiting signal =", sig)
}
