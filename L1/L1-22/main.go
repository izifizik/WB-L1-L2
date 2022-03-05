package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := 1<<20 + 500000
	b := 1 << 20

	fmt.Println("Div:", float64(a)/float64(b))
	fmt.Println("Mul:", a*b)
	fmt.Println("Sum:", a+b)
	fmt.Println("Sub:", a-b)

	a1 := big.NewFloat(2<<20 + 50000)
	b1 := big.NewFloat(2 << 20)
	sum := new(big.Float).Add(a1, b1)
	sub := new(big.Float).Sub(a1, b1)
	mul := new(big.Float).Mul(a1, b1)
	div := new(big.Float).Quo(a1, b1)
	fmt.Println("Mul:", mul)
	fmt.Println("Div:", div)
	fmt.Println("Sum:", sum)
	fmt.Println("Sub:", sub)
}
