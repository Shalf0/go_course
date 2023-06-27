package main

import (
	"flag"
	"fmt"
	"go-course/02-syntax/pkg/crawler"
	"go-course/02-syntax/pkg/crawler/spider"
	"log"
	"strings"
)

var links = []string{"https://go.dev/", "https://golang.org/"}

func main() {
	s := flag.String("s", "", "title filtering")
	flag.Parse()

	sp := spider.New()
	var dd []crawler.Document
	for _, v := range links {
		d, err := sp.Scan(v, 2)
		if err != nil {
			log.Fatal(err)
		}
		dd = append(dd, d...)
	}

	if *s != "" {
		for _, v := range dd {
			if strings.Contains(v.Title, *s) {
				fmt.Printf("URL: %s, TITLE: %s, BODY: %s\n", v.URL, v.Title, v.Body)
			}
		}
		return
	}
	for _, v := range dd {
		fmt.Printf("URL: %s, TITLE: %s, BODY: %s\n", v.URL, v.Title, v.Body)
	}
}
