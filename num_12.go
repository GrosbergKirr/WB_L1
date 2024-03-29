package main

import "fmt"

func Num12() {
	arr := []string{
		"cat", "cat", "dog", "cat", "tree", "dog", "cow",
	}
	var set []string
	mp := map[string]int{}

	for i := range arr {
		mp[arr[i]] += 1
	}

	for i := range mp {
		set = append(set, i)
	}

	fmt.Println(set)

}
