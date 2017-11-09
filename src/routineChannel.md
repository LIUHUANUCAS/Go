# Go 生产者消费者模型 

标签（空格分隔）： Go

---

## 1.生产者消费者编程模型

```go
func main() {
     in := make(chan int)
     out := make(chan int)
     go count(in)
     go squre(in, out)

     for x := range out { //读取管道，直到关闭，如果没有值，那么阻塞
          fmt.Println(x)
     }
}

func squre(in <-chan int, out chan<- int) {
     for x := range in { //读取管道，直到关闭管道为止，没有值，那么阻塞
          out <- x * x //写入管道
     }
     close(out) //关闭写入管道
}

func count(out chan<- int) {
     for i := 1; i < 10; i++ {
          out <- i //写入管道
     }
     close(out) //关闭写管道
}

```

## 2.channel的性质

- ### 1.使用make创建channel，且创建的变量为和，map,slice等一致均为引用类型

```go
    ch := make(chan int)
```

- ### 2.channel使用方式

- 2.1 向管道写入内容

```go
    x := 1
    ch <- x //写入内容
```

- 2.2 从管道读取内容

```go
    x := <- ch  //读取内容
```