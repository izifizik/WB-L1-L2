package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
)

//range for chan will end when chan will close down by someone
func main() {
	var c int
	fmt.Println("Set count of workers: ")
	_, err := fmt.Scan(&c)
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	m := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	for i := 0; i < c; i++ {
		wg.Add(1)
		go worker(m, wg, i)
	}

	go job(ctx, m)

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, os.Interrupt)
	sig := <-sign
	log.Printf("Caught signal %s. Shutting down...", sig)
	cancel()

	wg.Wait()
}

func worker(ch chan int, wg *sync.WaitGroup, id int) {
	for n := range ch {
		fmt.Println(id+1, ":", n)
	}
	fmt.Println("worker", id+1, "quit")
	wg.Done()
}

func job(ctx context.Context, ch chan int) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			close(ch)
			break
		default:
			ch <- i
			i++
			continue
		}
		break
	}
}
