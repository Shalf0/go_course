package index

import (
	"strings"
)

func Add(s string, id int, i map[string][]int) {
	for _, word := range strings.Split(s, " ") {
		if dd, ok := i[word]; ok {
			if !containInt(dd, id) {
				i[word] = append(i[word], id)
			}
			continue
		}
		i[word] = append(i[word], id)
	}
}

func containInt(dd []int, d int) bool {
	for _, v := range dd {
		if v == d {
			return true
		}
	}
	return false
}
