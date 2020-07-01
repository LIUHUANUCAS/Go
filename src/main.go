package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndSvr(":8080", nil)
}
