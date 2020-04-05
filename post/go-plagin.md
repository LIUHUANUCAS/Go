# Go 插件

## 概念

## 使用方法

- 插件定义


```golang
// Add.go
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

var Num Number

func Hello(s string) string {
	return fmt.Sprintf("hello %s,%+v,%d", s, Num, Num.Add())
}
func main() {}
```
- 编译方法：当前目录下`go build --buildmode=plugin Add.go`
- 会生成`Add.so`的动态链接库

- 插件使用

```golang
// main.go
package main

import (
	"log"
	"plugin"
)

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
}
```
- 运行方法: `go run main.go` 
- 需要把`Add.so` 放在当前目录下或者在代码(fname)里指定目录

## 用途

- 用于多个main函数运行

## 原理

- 本质的实现方法是使用`dlopen()`方法，所以plugin方式支持有目前有maxos,FreeBSD,linux等系统，不适用于Windows系统

## 参考内容
- [golang-plugin][1]

[1]:https://golang.org/pkg/plugin/