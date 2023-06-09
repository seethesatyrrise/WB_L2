package gettime

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

// Программа должна корректно обрабатывать ошибки библиотеки: выводить их в STDERR и возвращать ненулевой код выхода в OS
func GetTime() int {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return 1
	}
	fmt.Println(time)
	return 0
}
