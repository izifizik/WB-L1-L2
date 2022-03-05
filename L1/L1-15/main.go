package main

import "fmt"

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func someFunc2(size int) string {
	v := createHugeString(size)
	return v
}

func main() {
	someFunc()
	/*----------------------------------*/
	size := 1
	myString := someFunc2(size)
	fmt.Println(myString, len(myString))
}
func createHugeString(size int) (res string) {
	for i := 0; i < size; i++ {
		res += "A"
	}
	return
}
