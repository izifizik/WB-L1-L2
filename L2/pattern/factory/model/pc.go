package model

import (
	"fmt"
)

type pc struct {
	typeComputer string
	core         int
	ram          int
	monitor      int
}

func NewPC() Computer {
	return &pc{
		typeComputer: PCType,
		core:         16,
		ram:          32,
		monitor:      2,
	}
}

func (s pc) GetType() string {
	return s.typeComputer
}

func (s pc) Print() {
	fmt.Printf("%s: Core:%d, RAM:%d, Monitor:%d\n", s.typeComputer, s.core, s.ram, s.monitor)
}
