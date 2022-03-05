package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	//input
	s := `e\4\5`
	//call update
	ss, err := update(s)
	//check error
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(ss)
}

func update(s string) (string, error) {
	//"" => ""
	if s == "" {
		return "", nil
	}

	res := ""
	//v - previous char
	v := ""
	//count of "\"
	cnt := 0
	for _, c := range s {
		//checking for "\" and the fact that it wasn't repeated
		if v == "\\" && cnt == 1 {
			v = string(c)
			cnt = 0
			res += v
			continue
		}
		//checking for int
		n, err := strconv.Atoi(string(c))
		if err != nil {
			//checking for "\"
			if string(c) == "\\" {
				v = string(c)
				cnt++
				continue
			}
			//here default add all char's
			v = string(c)
			res += v
			continue
		}
		//repeat last char (v) n times
		for i := 0; i < n-1; i++ {
			res += v
		}
		//zeroing v because already has been used
		v = ""
	}
	//"45" => "" (некорректная строка)
	if res == "" {
		return "", fmt.Errorf("incorrect input")
	}
	//"a4bc2d5e" => "aaaabccddddde"
	//"abcd" => "abcd"
	return res, nil
}
