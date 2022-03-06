package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	t := time.Now().Add(time.Hour)
	d := time.Hour * 24 * 30

	if t.After(t1) && t.Before(t1.Add(d)) {
		fmt.Println(t)
		return
	}

	return
}
