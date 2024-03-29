package main

import "fmt"

func Num3() {
	arr := []int{2, 4, 6, 8, 10}
	var sm = 0
	ch := make(chan int, 5)
	defer close(ch)
	for i := range arr {
		go func() {
			ch <- arr[i] * arr[i]
		}()
	}
	for i := range arr {
		_ = i
		sm += <-ch
	}
	fmt.Println(sm)
}
