package main

import (
	"dev-11/internal/app"
	"log"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
