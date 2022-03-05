package main

import (
	"fmt"
	"sync"
)

//use WaitGroup for waiting end of all operations in goroutines
func main() {
	var arr = []int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}

	for _, c := range arr {
		wg.Add(1)
		go square(c, wg)
	}
	wg.Wait()
}

func square(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("sqrt(%d) : %d\n", n, n*n)
}
