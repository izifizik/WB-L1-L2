package main

import "fmt"

func main() {
	var strs = []string{"cat", "cat", "dog", "cat", "tree"}
	var res []string
	m := make(map[string]struct{})

	for _, c := range strs {
		_, ok := m[c]
		if ok {
			continue
		}

		m[c] = struct{}{}
		res = append(res, c)
	}
	fmt.Println(res)
}
