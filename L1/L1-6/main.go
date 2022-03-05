package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	c := context.Background()
	t := time.Second * time.Duration(5)

	cc, cancel := context.WithCancel(c)
	ct, cancel1 := context.WithTimeout(c, t)
	ch := make(chan struct{})

	fmt.Println("goroutine start by cancel")
	go contextCancel(cc)
	fmt.Println("goroutine start by timeout")
	go contextTimeout(ct)
	fmt.Println("goroutine start by chan")
	go chanel(ch)

	select {
	case <-time.After(t):
		cancel()
		cancel1()
		ch <- struct{}{}
	}
	time.Sleep(time.Millisecond * 100)
}

func contextCancel(c context.Context) {
	for {
		select {
		case <-c.Done():
			fmt.Println("goroutine stop by cancel")
			time.Sleep(time.Millisecond * 10)
			return
		}
	}
}

func contextTimeout(c context.Context) {
	for {
		select {
		case <-c.Done():
			fmt.Println("goroutine stop by timeout")
			time.Sleep(time.Millisecond * 10)
			return
		}
	}
}

func chanel(c chan struct{}) {
	for {
		select {
		case <-c:
			fmt.Println("goroutine stop by chanel")
			time.Sleep(time.Millisecond * 10)
			return
		}
	}
}
