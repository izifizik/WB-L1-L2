package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = i
	}
	go writeFromArr(ch1, arr)
	go sqrtFromChan(ch1, ch2)

	for c := range ch2 {
		fmt.Println(c)
	}
}

//writeFromArr write from arr to ch1
func writeFromArr(ch chan int, arr []int) {
	for _, c := range arr {
		ch <- c
	}
	close(ch)
}

//sqrtFromChan write value * 2 from ch1 to ch2
func sqrtFromChan(ch1 chan int, ch2 chan int) {
	for c := range ch1 {
		ch2 <- c * 2
	}
	close(ch2)
}
