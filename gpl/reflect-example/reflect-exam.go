package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	var x float64 = 1.9

	v := reflect.ValueOf(&x)
	log.Printf("cat set %t, type:%s %t", v.CanSet(), v.Type(), v.Elem().CanSet())
	v1 := v.Elem()
	v1.SetFloat(1.3)
	fmt.Println("hello ", x, v1.Interface())
}
