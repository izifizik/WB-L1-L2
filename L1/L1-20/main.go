package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "qwerty ahjgsdf jhasgdjhasg djh"
	fmt.Println(data + " - " + swapWords(data))
}

func swapWords(str string) string {
	split := strings.Split(str, " ")
	res := ""
	for _, s := range split {
		res = s + " " + res
	}
	return res
}
