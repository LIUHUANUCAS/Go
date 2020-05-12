# Go routine 内存实验

## Go routine 介绍

- 直接使用`go`关键字就可以创建一个routine
- 非常轻量级，内存占用较少

## 测试方法

- 开启大量go routine 并存在于内存中。
- 检测所有go routine所占用的内存
- go routine泄漏，阻塞go routine的完成

```golang
func main(){
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats // 获取内存字节数
		runtime.ReadMemStats(&s)
		return s.Sys
	}
	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() {
		wg.Done()// 增加计数，通知主携程完成标记
		<-c// 阻塞go routine ，造成go routine 泄漏，占用内存
	}

	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed() // 内存占用
	for i := 0; i < numGoroutines; i++ {
		go noop() // 开启go routine
	}
	wg.Wait()// 等待所有routine 完成
	after := memConsumed() // 内存占用
	// 计算内存开销
	fmt.Printf("%.3fkb\n", float64(after-before)/numGoroutines/1000)
}
```

- 获取内存占用情况`memConsumed`
- 阻塞go routine `<-c` ， 往未初始化的channel里放东西，会被阻塞
- 计算内存开销
- 结果：$numGoroutines=10^5$，占用内存大小`2.120kb`
- $numGoroutines=10^6$，占用内存大小`2.584kb`
- $numGoroutines=10^7$，占用内存大小`2.575kb`
- 可以看出内存开销非常小只有几kb大小

## 参考内容
- [go并发编程之道第三章][1]
- [go routine][2]

[1]:https://golang.org/pkg/plugin/
[2]:https://gobyexample.com/goroutines