package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	c, cl := context.WithCancel(context.Background())
	ch := make(chan int)
	defer close(ch)

	for i := 0; i < 100000000; i++ {
		go cCounter(c, ch)
	}
	ch <- 0

	select {
	case <-time.After(time.Second):
		cl()
		fmt.Println(<-ch)
	}
}

func cCounter(c context.Context, ch chan int) {
	for {
		select {
		case <-c.Done():
			return
		default:
			num := <-ch
			ch <- num + 1
		}
	}
}
