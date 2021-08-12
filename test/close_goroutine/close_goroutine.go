package main 

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	done := make(chan struct{})

	start := func(done <- chan struct{} ) <- chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			fmt.Println("start goroutine")
			for {
				select {
				case <- done:
					return 
				case ch <- rand.Int():
				}
			}
		}()
		return ch 
	}
	N := 10
	ch := start(done)
	for i:=0; i < N ; i++ {
		fmt.Println("data:", <- ch)
	}
	go func(){
		time.Sleep(time.Second*1)
		fmt.Println("terminated   start func")
		close(done)
	}()
	time.Sleep(time.Second*2)

}