package main

import "fmt"

func main() {
	data := "qwerty"
	fmt.Println(data + " - reverse is - " + swapChar(data))
}

func swapChar(str string) string {
	res := ""
	for _, c := range str {
		res = string(c) + res
	}
	return res
}
