package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index for /
// curl localhost:8080/  -->> Welcome!
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// Hello for /hello/name path
// curl localhost:8080/hello/world -> hello, world!
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func Hello3(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s %s !\n", ps.ByName("name"), ps.ByName("first"))
}

func main() {
	addr := ":8080"

	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/hello/:name/:first", Hello3)

	log.Printf("listen %s ...", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
