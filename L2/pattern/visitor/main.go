package main

import (
	"fmt"
	"visitor/cmd"
)

/*
Добавление логики не добавляя ее в сам интерфейс, а реализовывать в стороннем интерфейсе
Единственное, что надо это добавить метод вызова стороннего интерфейса
*/

func main() {
	s := cmd.Square{A: 2, B: 5}
	c := cmd.Circle{R: 14}

	calc := &cmd.AreaCalculator{}

	s.Accept(calc)
	fmt.Printf("Square area: %.2f\n", calc.Area)
	c.Accept(calc)
	fmt.Printf("Circle area: %.2f\n", calc.Area)
}
