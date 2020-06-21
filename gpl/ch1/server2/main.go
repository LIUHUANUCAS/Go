package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

// go run server2/main.go
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	defer mu.Unlock()
	mu.Lock()
	count++
	path := r.URL.Path
	fmt.Println(path)
	fmt.Printf("path:%q\n", path)
	fmt.Fprintf(w, "url.path: %q\n", path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	defer mu.Unlock()
	mu.Lock()
	fmt.Fprintf(w, "Count:%d\n", count)
}
