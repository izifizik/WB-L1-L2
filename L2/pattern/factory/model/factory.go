package model

import (
	"fmt"
)

const (
	ServerType = "server"
	PCType     = "pc"
	Laptop     = "laptop"
)

func New(typeOf string) Computer {
	switch typeOf {
	case ServerType:
		return NewServer()
	case PCType:
		return NewPC()
	case Laptop:
		return NewLaptop()
	default:
		fmt.Printf("Factory don't have this type of computer: %s\n", typeOf)
		return nil
	}
}
