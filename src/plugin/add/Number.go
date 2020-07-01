package main

import (
	"fmt"
)

type Number struct {
	A, B int32
}

func (n *Number) String() string {
	return fmt.Sprintf("[%d,%d]", n.A, n.B)
}
func (n *Number) Add() int32 {
	return n.A + n.B
}
