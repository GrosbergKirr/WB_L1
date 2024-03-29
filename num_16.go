package main

import "fmt"

func Num16() {
	arr := []int{1, 2, 34, 5, 6, 56, 5675, 3435, 63, 335, 45, 3, 657, 5652, 53, 8}

	fmt.Println(quickSort(arr))
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivotIndex := len(arr) / 2
	pivot := arr[len(arr)/2]

	left, right := partition(arr, pivotIndex)

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}

func partition(arr []int, pivotIndex int) ([]int, []int) {
	pivot := arr[pivotIndex]
	var left, right []int

	for i, v := range arr {
		if i == pivotIndex {
			continue
		}
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return left, right
}
