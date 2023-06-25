// Package curl.
// Its name means concurrent url, not curl itself
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		go fmt.Println(<-ch) // В оригинальном решении эта строчка кода запускалась без горутины и были дедлоки
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	secs := time.Since(start)
	ch <- fmt.Sprintf("%.2fs   %7d    %s", secs, b, url)
}
