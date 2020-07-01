package main

import (
 "fmt"
 "reflect"
 "io"
)
type Action struct {
	Name string
	ID int32
}
func (a *Action)Str()string {
	return a.Name
}
func (a Action)Str1()string{
	return a.Name
}
func GetName(a *Action) string{
	return a.Name
}
func main() {
a := Action{
	Name :"hello",
	ID:1,
}
 var x float64 = 3.4
 t := reflect.TypeOf(x)
 fmt.Println("type:", t)
 fmt.Println(t.Size())
  d := reflect.TypeOf(a)
  fmt.Println(d)
  fmt.Println(d.Method(0))
  fmt.Println(d.Kind())
  fmt.Println(d.PkgPath())
  fmt.Println(d.NumField())
//   f := reflect.ValueOf(GetName)
	// var c io.Writer
	var c  ByteCounter
	b := []byte("hello world")
	w,n := CountingWriter(&c)
	w.Write(b)
	fmt.Println(*n,len(b))
	b = []byte("hello ")
	w.Write(b)
	fmt.Println(*n,len(b))


}

type ByteCounter int
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}
type WriteCount struct {
	n int64
	w io.Writer
}
func(w *WriteCount)Write(b []byte)(int ,error ) {
	w.n += int64(len(b))
	return w.w.Write(b)
}
func CountingWriter(w io.Writer) (io.Writer, *int64){
	wc := &WriteCount{
		w :w,
		n :0,
	}
	return wc,&wc.n
}