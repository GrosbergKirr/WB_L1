package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func Num5() {

	var t int
	fmt.Println("Enter work time (in seconds): ")
	_, err := fmt.Scanln(&t)
	if err != nil {
		panic(err)
	}

	workTime := time.Second * time.Duration(t)
	ctx, _ := context.WithTimeout(context.Background(), workTime)

	ch := make(chan string)
	go ReadFromCh(ctx, ch)
	i := 0
loop:
	for {
		s := "Hello " + strconv.Itoa(i)
		select {
		case <-ctx.Done():
			break loop
		case ch <- s:
			i++
		}
	}
	fmt.Println("end")
}

func ReadFromCh(ctx context.Context, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case res := <-ch:
			fmt.Println(res)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
