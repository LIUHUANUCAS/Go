package main

import (
	"fmt"
)

var Num Number

func Hello(s string) string {
	return fmt.Sprintf("hello %s,%+v,%d", s, Num, Num.Add())
}

func main() {}
