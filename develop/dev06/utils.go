package main

import (
	"fmt"
)

// вывод описаний ключей
func getHelp() {
	fmt.Println("-f - \"fields\" - выбрать поля (колонки)\n" +
		"-d - \"delimiter\" - использовать другой разделитель\n" +
		"-s - \"separated\" - только строки с разделителем")
}
