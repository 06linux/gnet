package gbase

import (
	"fmt"
)

type baseTime struct{}

func (baseTime) GetTime() {
	fmt.Println("randUtil getTime")
}
