package main

import (
	// "bufio"
	"fmt"
	"io"
	// "io/ioutil"
	"net/http"
	"os"
	"strings"
)

// go run fetch.go
func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("resp.Status:%s\n", resp.Status)
		wlen, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%d err:%s", wlen, err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: reading %s:%v\n", url, err)
		// 	os.Exit(1)
		// }
		// fmt.Printf("%s", b)
		// io.Copy()
	}
}
