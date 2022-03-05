package main

import "fmt"

//Human struct:
type Human struct {
	age uint
}

//Action struct:
type Action struct {
	//inheritance Human struct
	Human
	name string
}

//HowOld - is a method for struct Human as a parent and Action as a child
func (h Human) HowOld() uint {
	return h.age
}

func main() {
	h := Human{age: 12}
	fmt.Printf("How old is human : %d\n", h.HowOld())

	act := Action{Human: h, name: "Harry"}
	fmt.Printf("How old is action %s : %d\n", act.name, act.HowOld())

	a := Action{name: "Billy"}
	a.age = 10
	fmt.Printf("How old is action %s : %d\n", a.name, a.HowOld())

}
