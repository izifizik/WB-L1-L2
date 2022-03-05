package main

import (
	"bufio"
	"flag"
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

func usage() {
	fmt.Println("Usage: [-k указание колонки для сортировки] [-n сортировать по числовому значению] " +
		"[-r сортировать в обратном порядке] [-u не выводить повторяющиеся строки]")
	os.Exit(0)
}

func main() {
	k := flag.Int("k", 0, "указание колонки для сортировки")
	n := flag.Bool("n", false, "сортировать по числовому значению")
	r := flag.Bool("r", false, "сортировать в обратном порядке")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")

	//M := flag.Bool("M", false, "сортировать по названию месяца")
	//b := flag.Bool("b", false, "игнорировать хвостовые пробелы")
	//c := flag.Bool("c", false, "проверять отсортированы ли данные")
	//h := flag.Bool("h", false, "сортировать по числовому значению с учетом суффиксов")

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		usage()
	}

	file, err := os.Open(args[0])
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()

	text := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("before:", text)
	fmt.Println("after :", task(text, *k-1, *n, *r, *u))
}

//task return sorted slice string
func task(arr []string, k int, n, r, u bool) []string {
	//if M {
	//	return month(arr)
	//}
	// конкретная строка для действа если в диапазоне строк в файле иначе все строки
	if k >= 0 && k < len(arr) {
		split := strings.Split(arr[k], " ")
		v := ""
		for _, s := range split {
			v += flagCheck(s, n, r)
		}
		arr[k] = v
		return arr
	}

	res := make([]string, 0)
	for _, str := range arr {
		split := strings.Split(str, " ")
		for _, s := range split {
			st := flagCheck(s, n, r)
			if st == "" {
				continue
			}
			res = append(res, st)
		}
	}

	if u {
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

//func isNum(arr []string) bool {
//	str := arr[0]
//	r := str[0]
//	if r > '0' && r < '9' {
//		return false
//	}
//	return true
//}
//func month(arr []string) []string {
//	ok := isNum(arr)
//	if ok {
//		sort.Strings(arr)
//	}
//	return arr
//}
