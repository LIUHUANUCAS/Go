package main

import (
	"fmt"
	"sync"
	"time"
)

/**
  main function
*/
const (
	N          = 1000000
	Buflen     = 1000
	concurency = 100
)

func main() {
	number := make([]int, 10*Buflen)
	for i := 1; i < 10*Buflen; i++ {
		number[i] = i
	}
	Concurency(number)
}

func Concurency(number []int) {

	start := time.Now()
	defer func() {
		cost := time.Since(start)
		fmt.Printf("muliti-goroutine with buffer channel cost:%s\n", cost)
	}()
	inputChan := make(chan int, len(number)) //消费缓冲区
	resChan := make(chan int, len(number))   //结果缓冲区

	wg := sync.WaitGroup{}

	for i := 0; i < concurency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for x := range inputChan { //设置的并发线程从消费缓冲区里拿数据
				y, err := DealFunction(x) //并发的处理函数
				if err != nil {
					fmt.Printf("error err:%s x:%d", err.Error(), x)
					return
				}
				resChan <- y // 结果回写缓冲区
			}
		}()
	}

	for _, x := range number { // 向消费缓冲区放数据
		inputChan <- x
	}
	close(inputChan)

	wg.Wait()
	close(resChan)

	/**deal with the result */
	for x := range resChan { // 处理结果
		// fmt.Printf("x=%d\n", x)
		x += 1
	}
}

func DealFunction(x int) (int, error) { //处理函数
	return x + 1, nil
}
