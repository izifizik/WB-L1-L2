package main

import (
	"factory/model"
	"fmt"
	"log"
)

/*
Фабрика дает легкость взаимодействия с интерфейсом
Все что нужно вызвать нужный конструктор и все сделается за тебя
*/

func main() {
	server := model.New(model.ServerType)
	if server == nil {
		log.Fatal()
	}
	fmt.Println(server.GetType())
	server.Print()

	pc := model.New(model.PCType)
	fmt.Println(pc.GetType())
	pc.Print()

	laptop := model.New(model.Laptop)
	if laptop == nil {
		log.Fatal()
	}
	fmt.Println(laptop.GetType())
	laptop.Print()

	mono := model.New("monoblok")
	if mono == nil {
		log.Fatal()
	}
	fmt.Println(mono.GetType())
	mono.Print()
}
