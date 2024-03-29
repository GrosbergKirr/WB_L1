package main

import "fmt"

func Num23() {
	arr := []int{1, 2, 34, 5, 6, 56, 5675, 3435, 63, 335, 45, 3, 657, 5652, 53, 8}

	i := 2

	copy(arr[i:], arr[i+1:])

	fmt.Println(arr)

	arr = arr[:len(arr)-1]

	fmt.Println(arr)
}
