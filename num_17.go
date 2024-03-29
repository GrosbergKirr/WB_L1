package main

import "fmt"

func Num17() {
	arr := []int{1, 3, 5, 6, 9, 11, 45, 78, 90, 100, 103, 105, 105, 109}
	fmt.Println(BinarySearch(arr, 5))
}
func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	ind := -1
	flag := -1
	for left <= right {
		//fmt.Println("-------------------")
		mid := (left + right) / 2
		//fmt.Println(arr[mid])
		if arr[mid] == target {
			ind = mid
			flag = ind
			break

		}
		if arr[mid] > target {
			right = mid - 1
			//fmt.Println("move left")
		} else {
			left = mid + 1
			ind += mid + 1
			//fmt.Println("move right")

		}
	}
	return flag
}
