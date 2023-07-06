package index

import (
	"strings"
)

// Update adds unique id to key in a provided index.
// If id is present in key, no value is added.
func Update(s string, id int, idx map[string][]int) {
	for _, word := range strings.Split(s, " ") {
		if dd, ok := idx[word]; ok {
			if containInt(dd, id) {
				continue
			}
		}

		idx[word] = append(idx[word], id)
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
