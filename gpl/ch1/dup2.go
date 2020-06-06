package main

import (
	"bufio"
	"fmt"
	"os"
)

// go run dup2.go < echo1.go
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countlines(os.Stdin, counts)
	} else {
		for _, fname := range files {
			f, err := os.Open(fname)
			if err != nil {
				fmt.Fprintf(os.Stderr, "err:%s", err)
				continue
			}
			countlines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
func countlines(f *os.File, countmap map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		countmap[input.Text()]++
	}
}
