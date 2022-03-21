package main

import (
	"builder/model"
)

/*
Строитель дает инкапсуляцию создания объекта
*/

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
