package main

import (
	"fmt"
	"time"
)

/**
  main function
*/
const (
	N = 1000000
	Buflen = 100
)

func main(){
	main_channel()
	main_channel_buffer()
}

func main_channel_buffer() {
	start := time.Now()

	defer func(){
		cost := time.Since(start)
		fmt.Println("channel_buffer_cost=",cost)
	}()

	in := make(chan int,Buflen)
	out := make(chan int,Buflen)
	go Producer(in)
	go Consumer(in, out)

	for x := range out {
		x += 1
	}

}

func main_channel() {
	start := time.Now()
	defer func(){
		cost := time.Since(start)
		fmt.Println("channel_cost=",cost)
	}()
	in := make(chan int)
	out := make(chan int)
	go Producer(in)
	go Consumer(in, out)

	for x := range out {
		x += 1
	}
}

func Consumer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func Producer(out chan<- int) {
	for i := 1; i < N; i++ {
		out <- i
	}
	close(out)
}



















