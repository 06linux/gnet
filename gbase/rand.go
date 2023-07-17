package gbase

import "fmt"

type baseRand struct{}

func (baseRand) GetRand() {
	fmt.Println("baseRand getRand")
}
