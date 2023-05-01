package controller

import "fmt"

type baseCtrl struct {
}

func (b *baseCtrl) Message() {
	fmt.Println("baseCtrl")
}

var (
	BaseCtrl = new(baseCtrl)
)
