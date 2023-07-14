package gutil

import (
	"fmt"
)

type timeUtil struct{}

func (timeUtil) GetTime() {
	fmt.Println("randUtil getTime")
}
