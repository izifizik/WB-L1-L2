package main

import "fmt"

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int][]float64)

	for _, c := range arr {
		gap := (int(c) / 10) * 10
		m[gap] = append(m[gap], c)
	}
	fmt.Println(m)
}
