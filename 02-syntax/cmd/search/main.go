package main

import (
	"flag"
	"fmt"
	"go-course/02-syntax/pkg/crawler"
	"go-course/02-syntax/pkg/crawler/spider"
	"log"
	"strings"
)

// service - search service.
type service struct {
	s     *string
	sp    *spider.Spider
	links []string
}

// buildService constructs search service.
func buildService() *service {
	s := flag.String("s", "", "filters titles, prints all results if empty.")
	flag.Parse()

	sp := spider.New()

	return &service{
		s:     s,
		sp:    sp,
		links: []string{"https://go.dev/", "https://golang.org/"},
	}
}

func main() {
	app := buildService()

	var dd []crawler.Document
	for _, v := range app.links {
		d, err := app.sp.Scan(v, 2)
		if err != nil {
			log.Printf("Scan() err: %v", err)
			continue
		}
		dd = append(dd, d...)
	}

	if *app.s != "" {
		for _, v := range dd {
			if strings.Contains(v.Title, *app.s) {
				fmt.Printf("URL: %s, TITLE: %s, BODY: %s\n", v.URL, v.Title, v.Body)
			}
		}
		return
	}
	for _, v := range dd {
		fmt.Printf("URL: %s, TITLE: %s, BODY: %s\n", v.URL, v.Title, v.Body)
	}
}
