package gbase

import (
	"fmt"
)

type baseTime struct{}

var Time = &baseTime{}

func (baseTime) GetTime() {
	fmt.Println("randUtil getTime")
}
