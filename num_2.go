package main

import (
	"fmt"
	"os"
	"sync"
)

func Num2() {
	arr := []int{2, 4, 6, 8, 10}

	Myfunc(arr)
}
func Myfunc(arr []int) {

	wg := sync.WaitGroup{}

	for i := range arr {
		wg.Add(1)
		go func(arr []int) {
			fmt.Fprintln(os.Stdout, arr[i]*arr[i])
			wg.Done()
		}(arr)
	}

	wg.Wait()

}
