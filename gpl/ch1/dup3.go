package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// go run dup3.go
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	for _, fname := range files {
		data, err := ioutil.ReadFile(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err:%s", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
