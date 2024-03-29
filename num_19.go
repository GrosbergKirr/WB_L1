package main

import "fmt"

func Num19() {
	s := "mystring"
	r := []rune(s)

	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	fmt.Println(string(r))
}
