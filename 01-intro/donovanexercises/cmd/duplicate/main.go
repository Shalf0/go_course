package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	duplicate()
}

func duplicate() {
	counts := make(map[string]int)
	if len(os.Args) == 1 {
		countLines(os.Stdin, counts)
		for k, v := range counts {
			fmt.Printf("key: %s, num: %d\n", k, v)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	//t := time.NewTimer(10 * time.Second)
	for input.Scan() {
		counts[input.Text()]++
	}
}
