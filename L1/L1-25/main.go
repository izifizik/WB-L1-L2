package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("sleep")
	Sleep(time.Second)
	fmt.Println("after sleep")
}

func Sleep(d time.Duration) {
	select {
	case <-time.After(d):
		return
	}
}
