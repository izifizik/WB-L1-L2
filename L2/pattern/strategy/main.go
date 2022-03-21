package main

import (
	"strategy/navigator"
)

/*
Стратегия позволяет иметь разное поведение и выбирать реализацию от нужды логики
*/
func main() {
	var (
		start      = 10
		stop       = 100
		strategies = []navigator.Strategy{navigator.CarStrategy{}, navigator.SubwayStrategy{}, navigator.WalkStrategy{}}
	)
	nav := navigator.Navigator{}

	for _, strategy := range strategies {
		nav.SetStrategy(strategy)
		nav.Route(start, stop)
	}
}
