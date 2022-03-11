package navigator

import "fmt"

type Strategy interface {
	Route(start, stop int)
}

type CarStrategy struct {
}

func (s CarStrategy) Route(start, stop int) {
	avgSpeed := 30
	total := stop - start
	jam := 2
	time := (avgSpeed * 5 * total * jam) / 100
	fmt.Printf("Strategy: Car, Total: %d, Time: %dm\n", total, time)
}

type SubwayStrategy struct {
}

func (s SubwayStrategy) Route(start, stop int) {
	avgSpeed := 100
	total := stop - start
	waitTime := 5 + 5 + 2*5
	time := (avgSpeed*5*total + waitTime) / 100
	fmt.Printf("Strategy: Subway, Total: %d, Time: %dm\n", total, time)
}

type WalkStrategy struct {
}

func (s WalkStrategy) Route(start, stop int) {
	avgSpeed := 10
	total := stop - start
	time := avgSpeed * 10 * total / 100
	fmt.Printf("Strategy: Walk, Total: %d, Time: %dm\n", total, time)
}
