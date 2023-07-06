package main

import (
	"bufio"
	"fmt"
	"go-course/03-algorithms/pkg/crawler"
	"go-course/03-algorithms/pkg/crawler/spider"
	"log"
	"os"
	"sort"
)

var links = []string{"https://go.dev/", "https://golang.org/"}

func main() {
	sp := spider.New()

	var rr []crawler.Resource
	fmt.Println("Please wait for pages to be scanned.")
	for _, v := range links {
		l, err := sp.Scan(v, 2)
		if err != nil {
			log.Printf("Scan() err: %v", err)
			continue
		}
		sort.Slice(l.Docs, func(i, j int) bool {
			return l.Docs[i].ID < l.Docs[j].ID
		})
		rr = append(rr, l)
	}

	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Write a word to search")
	for input.Scan() {
		for _, r := range rr {
			for _, v := range r.Idx[input.Text()] {
				fmt.Println(binarySearch(r.Docs, v))
			}
		}
	}
}

func binarySearch(docs []crawler.Document, id int) string {
	var smallest, greatest = docs[0].ID, docs[len(docs)-1].ID

	for smallest <= greatest {
		mid := (smallest + greatest) / 2
		if docs[mid].ID == id {
			return docs[mid].URL
		}
		if docs[mid].ID < id {
			smallest = mid + 1
		}
		if docs[mid].ID > id {
			greatest = mid - 1
		}
	}

	return ""
}
