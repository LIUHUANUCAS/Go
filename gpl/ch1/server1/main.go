package main

import (
	"fmt"
	"log"
	"net/http"
)

// go run server1/main.go
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Println(path)
	fmt.Printf("path:%q\n", path)
	fmt.Fprintf(w, "url.path: %q\n", path)
}
