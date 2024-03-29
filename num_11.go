package main

import "fmt"

func Num11() {
	set1 := []int{1, 9, 2, 8, 3, 7, 4, 6, 5}
	set2 := []int{1, 7, 3, 8, 5, 12}

	mp := map[int]int{}

	for i := range set1 {
		mp[set1[i]] += 1
	}

	var set3 []int

	for i := range set2 {
		_, ok := mp[set2[i]]
		if ok {
			set3 = append(set3, set2[i])
		}

	}
	fmt.Println(set3)
}
