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

	var ll []spider.Link
	for _, v := range links {
		l, err := sp.Scan(v, 2)
		if err != nil {
			log.Printf("Scan() err: %v", err)
			continue
		}
		sort.Slice(l.Docs, func(i, j int) bool {
			return l.Docs[i].ID < l.Docs[j].ID
		})
		ll = append(ll, l)
	}

	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Please provide a word for search")
	for input.Scan() {
		for _, v := range ll {
			for _, word := range v.Idx[input.Text()] {
				fmt.Println(binarySearch(v.Docs, word))
			}
		}
	}
}

func binarySearch(docs []crawler.Document, d int) string {
	var smallest, greatest = docs[0].ID, docs[len(docs)-1].ID

	for smallest <= greatest {
		mid := (smallest + greatest) / 2
		if docs[mid].ID == d {
			return docs[mid].URL
		}
		if docs[mid].ID < d {
			smallest = mid + 1
		}
		if docs[mid].ID > d {
			greatest = mid - 1
		}
	}

	return ""
}
