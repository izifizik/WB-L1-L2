package main

import (
	"fmt"
)

//use chan for communication between goroutines
func main() {
	var sum int
	var arr = []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(arr))
	for _, n := range arr {
		go sqrt(n, ch)
	}

	for i := 0; i < len(arr); i++ {
		sum += <-ch
	}
	fmt.Println(sum)
}

func sqrt(n int, ch chan int) {
	ch <- n * n
}
