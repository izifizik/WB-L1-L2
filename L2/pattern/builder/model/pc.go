package model

import "fmt"

type PC struct {
	Core  int
	Brand string
	RAM   int
	GPU   int
}

func (p PC) Print() {
	fmt.Printf("%s: Core: %d, RAM: %d, GPU: %d\n", p.Brand, p.Core, p.RAM, p.GPU)
}
