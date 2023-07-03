// Package spider реализует сканер содержимого веб-сайтов.
// Пакет позволяет получить список ссылок и заголовков страниц внутри веб-сайта по его WebResource.
package spider

import (
	"go-course/03-algorithms/pkg/crawler"
	"go-course/03-algorithms/pkg/crawler/index"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Spider - служба поискового робота.
type Spider struct{}

// New - констрктор службы поискового робота.
func New() *Spider {
	return &Spider{}
}

// Scan осуществляет рекурсивный обход ссылок сайта, указанного в url,
// с учётом глубины перехода по ссылкам, переданной в depth.
func (s *Spider) Scan(url string, depth int) (crawler.Resource, error) {
	pages := make(map[string]string)
	err := parse(url, url, depth, pages)
	if err != nil {
		return crawler.Resource{}, err
	}

	res := crawler.Resource{
		Name: url,
		Docs: []crawler.Document{},
		Idx:  make(map[string][]int),
	}

	id := 0
	for url, title := range pages {
		res.Docs = append(res.Docs, crawler.Document{
			ID:    id,
			URL:   url,
			Title: title,
		})
		index.Update(title, id, res.Idx)

		id++
	}

	return res, nil
}

// parse рекурсивно обходит ссылки на странице, переданной в url.
// Глубина рекурсии задаётся в depth.
// Каждая найденная ссылка записывается в ассоциативный массив
// data вместе с названием страницы.
func parse(url, baseurl string, depth int, data map[string]string) error {
	if depth == 0 {
		return nil
	}

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	page, err := html.Parse(response.Body)
	if err != nil {
		return err
	}

	data[url] = pageTitle(page)

	if depth == 1 {
		return nil
	}
	links := pageLinks(nil, page)
	for _, link := range links {
		link = strings.TrimSuffix(link, "/")
		// относительная ссылка
		if strings.HasPrefix(link, "/") && len(link) > 1 {
			link = baseurl + link
		}
		// ссылка уже отсканирована
		if data[link] != "" {
			continue
		}
		// ссылка содержит базовый url полностью
		if strings.HasPrefix(link, baseurl) {
			parse(link, baseurl, depth-1, data)
		}
	}

	return nil
}

// pageTitle осуществляет рекурсивный обход HTML-страницы и возвращает значение элемента <tittle>.
func pageTitle(n *html.Node) string {
	var title string
	if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title = pageTitle(c)
		if title != "" {
			break
		}
	}
	return title
}

// pageLinks рекурсивно сканирует узлы HTML-страницы и возвращает все найденные ссылки без дубликатов.
func pageLinks(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				if !sliceContains(links, a.Val) {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = pageLinks(links, c)
	}
	return links
}

// sliceContains возвращает true если массив содержит переданное значение
func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
