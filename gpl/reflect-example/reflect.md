# 反射

## 使用方法

```go
var x float64 = 1.9
v := reflect.ValueOf(&x)
v1 := reflect.ValueOf(&x).Elem()

```
