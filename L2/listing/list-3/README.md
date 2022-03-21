Что выведет программа? Объяснить вывод программы.
```
package main

import (
	"fmt"
	"os"
)
func Foo() error {
	var err *os.PathError = nil
	return err
}
func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```
Ответ:
```
var err *os.PathError инициализация err типом os.PathError
так что при сравнении переменная err будет не nil a (*os.PathEerror).<nil> (где то видел что это так пишется но не уверен в точности записи)
```