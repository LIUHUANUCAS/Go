package main

import (
	"fmt"
)

var generator = func(done <-chan struct{}, numbers ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, v := range numbers {
			select {
			case <-done:
				return
			case ch <- v:
			}
		}
	}()
	return ch
}
var multiply = func(done <-chan struct{}, numStream <-chan int, factor int) <-chan int {
	next := make(chan int)
	go func() {
		defer close(next)
		for e := range numStream {
			select {
			case <-done:
				return
			case next <- e * factor:
			}
		}
	}()
	return next
}

var add = func(done <-chan struct{}, numStream <-chan int, addtive int) <-chan int {
	next := make(chan int)
	go func() {
		defer close(next)
		for v := range numStream {
			select {
			case <-done:
				return
			case next <- v + addtive:
			}
		}
	}()
	return next
}

func main() {
	done := make(chan struct{})
	defer close(done)
	ints := []int{1, 2, 3, 4}
	intStream := generator(done, ints...)
	pipeline := add(done, multiply(done, intStream, 10), 1)
	for v := range pipeline {
		fmt.Printf("%d ", v)
	}
}
