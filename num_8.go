package main

import (
	"strconv"
)

func Num8() {
	n := 77
	s := Bin(n)
	b := 3
	m := ""
	for i := len(s); i >= 1; i-- {
		if i == b {
			m += "1"
		} else {
			m += "0"
		}
	}
	//runes := []rune(s)

	//if string(runes[b]) == "0"{
	//	runes[b] =
	//}else{
	//	string(runes[b]) = "0"
	//}
	//fmt.Println(Bin(17))
	//m := 1 << 2
	//fmt.Println(m)
	//fmt.Printf("%b", m)
}

func Bin(n int) string {
	s := ""

	for n > 0 {
		s += strconv.Itoa(n % 2)
		n = (n - (n % 2)) / 2
	}
	res := ""
	for _, v := range s {
		res = string(v) + res
	}
	return res
}
