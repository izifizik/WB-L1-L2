package main

import (
	"factory/model"
	"fmt"
	"log"
)

func main() {
	server := model.New(model.ServerType)
	fmt.Println(server.GetType())
	server.Print()

	pc := model.New(model.PCType)
	fmt.Println(pc.GetType())
	pc.Print()

	laptop := model.New(model.Laptop)
	fmt.Println(laptop.GetType())
	laptop.Print()

	mono := model.New("monoblok")
	if mono == nil {
		log.Fatal()
	}
	fmt.Println(mono.GetType())
	mono.Print()
}
