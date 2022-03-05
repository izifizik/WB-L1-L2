package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	arr := make([]int, 5)
	for i := 0; i < 5; i++ {
		arr[i] = rand.Intn(1000)
	}
	fmt.Println(arr)
	qSort(arr)
	fmt.Println(arr)
}

func qSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left := 0
	right := len(arr) - 1
	pivot := len(arr) / 2

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	qSort(arr[:left])
	qSort(arr[left+1:])
}
