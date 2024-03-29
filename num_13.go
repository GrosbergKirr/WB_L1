package main

import "fmt"

func Num13() {
	n := 1
	m := 2
	fmt.Println(n, m)

	n, m = m, n

	fmt.Println(n, m)
}
