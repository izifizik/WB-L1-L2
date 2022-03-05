package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	arr := make([]int, 5)
	for i := 0; i < 5; i++ {
		arr[i] = rand.Intn(3)
	}
	sort.Ints(arr)
	find := 2
	num := bFind(find, arr)
	if num == -1 {
		fmt.Println("there is no solution")
		return
	}
	fmt.Printf("arr[%d] = %d", num, arr[num])

}

//bFind обязательное условие, что бы числа были > 0
func bFind(b int, arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	l := 0
	s := len(arr) / 2
	h := len(arr) - 1

	for l <= h {
		if arr[s] == b {
			return s
		}
		if arr[s] < b {
			h = s - 1
			s /= 2
			continue
		}
		l = s + 1
		s = (l + h) / 2
	}
	return -1
}
