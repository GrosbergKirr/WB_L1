package main

import (
	"fmt"
	"strings"
)

func Num26() {
	str := []string{"abcd", "abCdefAaf", "aabcd", "", "Cc", "string"}
	mp := map[string]int{}

	for _, v := range str {
		f := true
		for _, e := range v {
			_, ok1 := mp[string(e)]
			_, ok2 := mp[strings.ToUpper(string(e))]
			_, ok3 := mp[strings.ToLower(string(e))]
			if ok1 || ok2 || ok3 {
				f = false
				break
			}
			mp[string(e)] += 1

		}
		fmt.Println(f)
		fmt.Println("----------")
	}

}
