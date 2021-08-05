package main

import (
	"sync"
	"testing"
)

// go test -bench=. -cpu=1 context_switch_test.go
func BenchmarkContextSwitch(b *testing.B) {
	begin := make(chan struct{})
	ch := make(chan struct{})
	wg := sync.WaitGroup{}

	producer := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			ch <- struct{}{}
		}
	}
	consumer := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-ch
		}
	}
	wg.Add(2)
	go consumer()
	go producer()
	b.StartTimer()
	close(begin)
	wg.Wait()
	// close(ch)
}
