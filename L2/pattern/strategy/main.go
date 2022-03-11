package main

import (
	"strategy/navigator"
)

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
