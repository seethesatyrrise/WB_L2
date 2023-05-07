Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
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
<nil>
false

error - интерфейс, переменная err содержит в себе тип *os.PathError и значение nil. Из-за наличия 
типа переменная не считается равной nil, но при выводе значения выводится само значение этой 
переменной - nil.

```
