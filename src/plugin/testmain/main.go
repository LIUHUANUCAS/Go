package main

import (
	"log"
	// "os"
	"plugin"
	// "plugin/add"
)

// type Number struct {
// 	A, B int32
// }

func main() {
	fname := "Add.so"
	p, err := plugin.Open(fname)
	if err != nil {
		log.Println(err)
		return
	}
	h, err := p.Lookup("Hello")
	if err != nil {
		log.Println(err)
		return
	}
	hf := h.(func(string) string)
	r := hf("liuhuan")
	log.Println(r)
	// n, err := p.Lookup("Num")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// num := n.(*Number)
	// num.A = 10
	r = hf("end")
	log.Println(r)
}
