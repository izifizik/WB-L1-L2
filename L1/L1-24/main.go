package main

import "fmt"

type Point struct {
	x, y int
}

func NewPoint(x, y int) Point {
	return Point{x: x, y: y}
}

func (p Point) distance() int {
	d := p.x - p.y
	if d < 0 {
		return p.y - p.x
	}
	return d
}

func main() {
	p := NewPoint(1, 10)
	fmt.Println(p.distance())
}
