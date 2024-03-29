package main

import "fmt"

func Num10() {
	mp := map[int][]float64{}
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	for i := range arr {
		mp[int(arr[i])/10*10] = append(mp[int(arr[i])/10*10], arr[i])
	}
	fmt.Println(mp)
}
