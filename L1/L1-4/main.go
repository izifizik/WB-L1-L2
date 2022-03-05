package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

//range for chan will end when chan will close down by someone
func main() {
	var c int
	fmt.Println("Set count of workers: ")
	_, err := fmt.Scan(&c)
	if err != nil {
		log.Fatal(err.Error())
	}

	m := make(chan int)
	defer close(m)

	for i := 0; i < c; i++ {
		go worker(m, i)
	}

	go job(m)

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, os.Interrupt)
	sig := <-sign
	log.Printf("Caught signal %s. Shutting down...", sig)
}

func worker(ch chan int, id int) {
	for n := range ch {
		fmt.Println(id+1, ":", n)
	}
}

func job(ch chan int) {
	i := 1
	for {
		ch <- i
		i++
		time.Sleep(time.Millisecond)
	}
}
