Что выведет программа? Объяснить вывод программы.
```
package main
import (
	"fmt"
)
func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}
func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}
func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```
Ответ:
```
Конструкция defer срабатывает после вызова return 
test() (x int) - функция возвращающая конкретную переменную x без разницы будет ли она в return или нет
anotherTest() int - вернет только то что будет в return
```