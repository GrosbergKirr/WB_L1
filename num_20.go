package main

import (
	"fmt"
	"strings"
)

func Num20() {
	s := "Grosberg Kirill Alex age25"
	arr := strings.Split(s, " ")
	var newS string
	for i := len(arr) - 1; i >= 0; i-- {
		newS += arr[i]
		newS += " "

	}
	fmt.Println(newS)
}
