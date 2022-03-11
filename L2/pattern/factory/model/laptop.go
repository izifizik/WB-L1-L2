package model

import (
	"fmt"
)

type laptop struct {
	typeComputer string
	core         int
	ram          int
}

func NewLaptop() Computer {
	return &laptop{
		typeComputer: Laptop,
		core:         2,
		ram:          4,
	}
}

func (s laptop) GetType() string {
	return s.typeComputer
}

func (s laptop) Print() {
	fmt.Printf("%s: Core:%d, RAM:%d\n", s.typeComputer, s.core, s.ram)
}
