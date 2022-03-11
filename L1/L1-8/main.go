package main

import (
	"fmt"
	"log"
	"strconv"
)

//Operator xor (^) if 0 -> 1 else 1 -> 0
func main() {
	var num int64 = 10234
	bitNum := 5
	if bitNum < 0 {
		log.Fatal("bit num must be > 0")
	}
	bitVal := strconv.FormatInt(num, 2)
	fmt.Println("binary:", bitVal)

	fmt.Println("----------------------After----------------------")
	bitNewVal := strconv.FormatInt(num^(1<<bitNum), 2)
	fmt.Println("binary:", bitNewVal)
}
