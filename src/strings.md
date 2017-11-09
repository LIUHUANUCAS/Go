# Go strings 包常用函数 

标签（空格分隔）： Go 

---

## 1.字符串分割和合并
- 1.1 字符串分割函数原型：
```go
    strings.Join(slist []string,sep string) string
```
- 1.2 字符串合并函数原型：
```go
    strings.Join(str string,sep string) []string
```

- 1.3 实验代码
```go
func main() {
	strlist := []string{"hello","world","Go"}
	str := strings.Join(strlist, ",")// 对列表按照某种分隔符进行合并
	fmt.Println(str)
	splitlist := strings.Split(str, ",") // 对字符串按照某种字符方式进行分割
	fmt.Println(splitlist)
}

```

## 2.字符串去除两边的空格
- 2.1 函数原型：
```go
    strings.TrimSpace(str string) string
```

- 2.2 实验代码
```go
func main() {
	intstr := " 1 123  34 "
	intstr = strings.TrimSpace(intstr)
	intlist := strings.Split(intstr, " ")
	fmt.Printf("list=%v", intlist)
	fmt.Printf("str=%s", strings.Join(intlist, ",")) //[1,123,,34]
	for _, v := range intlist {
		v = strings.TrimSpace(v) // 空格仍然存在于分割分割结果当中
		if v == "" {
			continue
		}
		i, _ := strconv.Atoi(v)
		fmt.Printf("v=[%d],i+1=[%d]", v, i+1)
	}
}
```
>- 可以看出来，对于字符串分割函数的使用，分割的结果是存在分割符号的，所以需要注意
>- 对于分割符号，我们还需要在判断是否存在与结果当中

## 3.字符串转数字

**由于Go当中数字类型种类很多，所以字符串转数字的函数也较多**

- 3.1 字符串转int
第一个函数是字符串转int类型，如果转换成功那么err==nil，返回的数字可用，否则转换不成功
```go
    strconv.Atoi(str string)(int,error)
```
