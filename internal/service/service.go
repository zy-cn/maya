package service

import "fmt"

type baseSrv struct {
}

func (b *baseSrv) Message() {
	fmt.Println("baseSrv")
}

var (
	BaseSrv = new(baseSrv)
)
