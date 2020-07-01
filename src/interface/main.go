package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	var c ByteCounter
	b := []byte("hello world")
	w, n := CountingWriter(&c)
	w.Write(b)
	fmt.Println(*n, len(b))
	b = []byte("hello ")
	w.Write(b)
	fmt.Println(*n, len(b))
	http.ListenAndSvr()
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

func (w *WriteCount) Write(b []byte) (int, error) {
	w.n += int64(len(b))
	return w.w.Write(b)
}

// Wr ite a function CountingWriter with the sig nature below that, given an
// io.Writer, retur nsanew Writer that wraps the original, and a point er to an int64 var iable
// that at any mom ent contains the number of bytes writt en to the new Writer.
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	wc := &WriteCount{
		w: w,
		n: 0,
	}
	return wc, &wc.n
}
