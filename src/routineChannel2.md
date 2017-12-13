# 并发模型的生产者消费者模型

标签（空格分隔）： Go

---

- 基本模型：带有缓冲区的生产者和消费者的并发结构
- 两个缓冲区：一个消费缓冲区，一个结果缓冲区，一个生产者，多个消费者线程
- 多个线程：每个线程从消费缓冲区里拿数据，把结果缓存在结果缓冲区中
- 以下面的例子说明，可控制线程数量的生产者消费者模型



```go
const (
	concurency = 100
)

func Concurency(number []int) {

	start := time.Now()
	defer func(){
		cost := time.Since(start)
		fmt.Printf("muliti-goroutine with buffer channel cost:%s\n",cost)
	}()

	inputChan := make(chan int, len(number)) //消费缓冲区
	resChan := make(chan int, len(number))	//结果缓冲区

	wg := sync.WaitGroup{}

	for i := 0; i < concurency; i++ { 

		wg.Add(1)
		go func() {

			defer wg.Done()

			for x := range inputChan { //设置的并发线程从消费缓冲区里拿数据
				y, err := DealFunction(x) //并发的处理函数
				// error handler
				resChan <- y // 结果回写缓冲区
			}

		}()
	}

	for _, x := range number { // 向消费缓冲区放数据
		inputChan <- x
	}

	close(inputChan)// 关闭消费缓冲区，通知消费者线程所有数据已经处理完
	wg.Wait()  // 等待所有消费者线程处理完所有的数据，结束线程

	close(resChan)// 关闭结果缓冲区，通知结果处理线程，这里是main线程

	/**deal with the result */
	for x := range resChan {  // 处理结果
		// fmt.Printf("x=%d\n", x)
		x += 1
	}
}

```

- 具体的处理函数

```go
func DealFunction(x int) (int, error) { //处理函数
	return x + 1, nil
}
```





