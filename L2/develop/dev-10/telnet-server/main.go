package main

import (
	"bufio"
	"context"
	"dev-10/telnet-server/config"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("Server start")
	cfg := config.NewConfig()
	l, err := net.Listen("tcp", cfg.Host+":"+cfg.Port)
	if err != nil {
		log.Fatalln("Internal server error: " + err.Error())
	}
	defer l.Close()
	conn, err := l.Accept()
	if err != nil {
		log.Fatalln("Internal server error: " + err.Error())
	}
	defer conn.Close()

	ctx, c := context.WithCancel(context.Background())
	defer c()
	input := make(chan string)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)

	go Read(ctx, input, conn)
	go Write(ctx, input, conn)

	<-sig
	log.Println("Server stop")
}

func Read(c context.Context, input chan string, conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		select {
		case <-c.Done():
			close(input)
			return
		default:
			msg, err := reader.ReadString('\n')
			if err == nil {
				input <- msg
			}
		}
	}
}

func Write(c context.Context, input chan string, conn net.Conn) {
	for {
		select {
		case <-c.Done():
			return
		case msg := <-input:
			fmt.Fprintf(conn, "Server: "+msg)
		}
	}
}
