package main

import (
	"bufio"
	"dev-10/config"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	cfg := config.NewConfig()

	timeout, err := strconv.Atoi(cfg.Timeout[:len(cfg.Timeout)-1])
	if err != nil {
		log.Fatalln("Client:" + err.Error())
	}
	to := time.Duration(timeout) * time.Second

	var conn net.Conn

	t := time.Now()
	for time.Since(t) < to {
		conn, err = net.Dial("tcp", cfg.Host+":"+cfg.Port)
		if err == nil {
			break
		}
	}
	if err != nil {
		log.Fatalln("Client:", err.Error())
	}
	defer conn.Close()

	log.Printf("Client: connected to %s:%s", cfg.Host, cfg.Port)
	go Read(conn)

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		input := scan.Text()
		_, err = fmt.Fprintf(conn, input+"\n")
		if err != nil {
			log.Fatalln("Client: server close connection")
		}
	}
}

func Read(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err == io.EOF {
			log.Fatalln("Client:", err.Error())
		}
		if err != nil {
			log.Println("Client:", err.Error())
		}
		fmt.Println("Client: message from server:", msg)
	}
}
