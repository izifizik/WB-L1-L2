package main

import (
	"fmt"
	"sync"
)

var m sync.Mutex
var mm map[int]string

func main() {
	n := "data"
	wg := &sync.WaitGroup{}
	mm = make(map[int]string)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go read(i, n, wg)
	}

	wg.Wait()
	for i, c := range mm {
		fmt.Println(i, ":", c)
	}
}

func read(id int, n string, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Lock()
	mm[id] = n
	m.Unlock()
}
