package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	//alternative
	//alternative()
	c, cl := context.WithCancel(context.Background())
	i := new(uint64)

	for j := 0; j < 100; j++ {
		go cCounterA(c, i)
	}

	select {
	case <-time.After(time.Second):
		cl()
		fmt.Println(*i)
	}
}

func cCounterA(c context.Context, i *uint64) {
	for {
		select {
		case <-c.Done():
			return
		default:
			atomic.AddUint64(i, 1)
		}
	}
}

func alternative() {
	c, cl := context.WithCancel(context.Background())
	ch := make(chan int)
	defer close(ch)

	for i := 0; i < 100; i++ {
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
