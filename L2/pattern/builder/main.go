package main

import (
	"builder/model"
)

func main() {
	asusCollector := model.GetCollector("ASUS")
	hpCollector := model.GetCollector("HP")

	factory := model.NewFactory(asusCollector)
	asusComputer := factory.CreateComputer()
	asusComputer.Print()

	factory.ChangeCollector(hpCollector)
	hpComputer := factory.CreateComputer()
	hpComputer.Print()
}
