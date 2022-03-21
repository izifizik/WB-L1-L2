package main

import (
	"bufio"
	"dev-3/config"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// for realisation "Interface" interface package sort
type sortString []rune

func (s sortString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortString) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s sortString) Len() int {
	return len(s)
}

func getFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var text []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	return text, nil
}

func main() {
	cfg := config.NewConfig()

	text, err := getFile(cfg.Filename)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("before:", text)
	fmt.Println("after :", task(text, cfg))
}

//task return sorted slice string
func task(arr []string, cfg config.Config) []string {
	//if M {
	//	return month(arr)
	//}
	// конкретная строка для действа если в диапазоне строк в файле иначе все строки
	if cfg.K >= 0 && cfg.K < len(arr) {
		split := strings.Split(arr[cfg.K], " ")
		v := ""
		for _, s := range split {
			v += flagCheck(s, cfg.N, cfg.R)
		}
		arr[cfg.K] = v
		return arr
	}

	res := make([]string, 0)
	for _, str := range arr {
		split := strings.Split(str, " ")
		for _, s := range split {
			st := flagCheck(s, cfg.N, cfg.R)
			if st == "" {
				continue
			}
			res = append(res, st)
		}
	}

	if cfg.U {
		return checkUnique(res)
	}

	return res
}

//flagCheck return string pass all of flags
func flagCheck(str string, n, r bool) string {
	if r {
		str = RSort(sortString(str))
	} else {
		str = Sort(sortString(str))
	}

	if n {
		str = number(str)
	}

	return str
}

//checkUnique - flag -u
func checkUnique(res []string) []string {
	uq := make([]string, 0)
	for i, str := range res {
		ok := false
		for j := i + 1; j < len(res); j++ {
			if str == res[j] {
				ok = true
				break
			}
		}
		if ok {
			continue
		}
		uq = append(uq, str)
	}
	return uq
}

// number - flag -n
func number(str string) string {
	res := ""
	for _, c := range str {
		if c < '0' || c > '9' {
			break
		}
		res += string(c)
	}
	return res
}

//RSort - flag -r
func RSort(str sortString) string {
	sort.Sort(sort.Reverse(str))
	return string(str)
}

//Sort - without flag -r
func Sort(str sortString) string {
	sort.Sort(str)
	return string(str)
}
