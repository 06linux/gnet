package gutil

import "fmt"

type randUtil struct{}

func (randUtil) GetRand() {
	fmt.Println("randUtil getRand")
}
