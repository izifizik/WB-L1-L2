package main

import (
	"Work/L2/pattern/builder/model"
)

/*
+:
Пошаговое создание чего-либо (из составляющих частей)
Код не повторяется
Все под интерфейсом
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
