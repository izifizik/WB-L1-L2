package model

import (
	"fmt"
)

type server struct {
	typeComputer string
	core         int
	ram          int
}

func NewServer() Computer {
	return &server{
		typeComputer: ServerType,
		core:         64,
		ram:          256,
	}
}

func (s server) GetType() string {
	return s.typeComputer
}

func (s server) Print() {
	fmt.Printf("%s: Core:%d, RAM:%d\n", s.typeComputer, s.core, s.ram)
}
