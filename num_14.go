package main

import "fmt"

func Num14() {

	var i interface{} = 2
	var s interface{} = "string"
	var ch interface{} = make(chan int)
	var arr []interface{}
	arr = append(arr, i)
	arr = append(arr, s)
	arr = append(arr, ch)
	Switcher(arr)
}

func Switcher(arr []interface{}) {

	for i := range arr {
		switch arr[i].(type) {
		case int:
			fmt.Println("is int")
		case string:
			fmt.Println("is string")
		case chan int:
			fmt.Println("is chan int")
		}
	}
}
