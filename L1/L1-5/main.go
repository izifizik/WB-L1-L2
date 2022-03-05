package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	N := 5
	read := make(chan struct{})
	write := make(chan struct{})
	data := make(chan []byte)

	go writeF(write, data)
	go readF(read, data)

	//select with 1 case will execute only it
	select {
	case <-time.After(time.Second * time.Duration(N)):
		write <- struct{}{}
		read <- struct{}{}
		os.Exit(1)
	}
}

func readF(ch chan struct{}, c chan []byte) {
	for {
		select {
		case <-ch:
			return
		default:
			log.Printf("%s", <-c)
		}
	}
}

func writeF(ch chan struct{}, c chan []byte) {
	i := 0
	for {
		select {
		case <-ch:
			return
		default:
			c <- []byte("data " + strconv.Itoa(i))
			i++
		}
	}
}
