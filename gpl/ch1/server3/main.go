package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

// go run server3/main.go
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	defer mu.Unlock()
	mu.Lock()
	count++
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "header:[%q]= %q\n", k, v)
	}
	fmt.Fprintf(w, "Host:%q\n", r.Host)
	fmt.Fprintf(w, "remoteAddr:%q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "form:[%q]=%q\n", k, v)
	}

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
