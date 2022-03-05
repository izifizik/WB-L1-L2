package main

import (
	"fmt"
	"reflect"
)

func main() {
	var val interface{}
	val = 1
	fmt.Printf("%T\n", val)
	val = "asd"
	fmt.Printf("%T\n", val)
	val = true
	fmt.Println(reflect.TypeOf(val))
	val = make(chan struct{})
	fmt.Println(reflect.TypeOf(val))

}
