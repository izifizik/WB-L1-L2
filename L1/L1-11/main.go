package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	set1 := make([]int, 10)
	set2 := make([]int, 5)

	for i := 0; i < 10; i++ {
		set1[i] = rand.Intn(50)
	}
	fmt.Println("set1:", set1)
	for i := 0; i < 5; i++ {
		set2[i] = rand.Intn(50)
	}
	fmt.Println("set2:", set2)

	fmt.Println(getIntersection(set1, set2))
}

func getIntersection(set1, set2 []int) (res []string) {
	m := make(map[int]struct{})
	for _, c := range set1 {
		m[c] = struct{}{}
	}

	for _, c := range set2 {
		_, ok := m[c]
		if !ok {
			delete(m, c)
			continue
		}
		res = append(res, strconv.Itoa(c))
	}

	return
}
