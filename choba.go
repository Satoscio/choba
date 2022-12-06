package choba

import (
	"fmt"
)

func Choba(n ChobaInt) {
	var i ChobaInt
	for i = 0; i < n; i++ {
		fmt.Print("Choba")
	}
	fmt.Println()
}

func ToChobaInt(n int) (x ChobaInt) {
	x = ChobaInt(n)
	return
}