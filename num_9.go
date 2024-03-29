package main

import (
	"fmt"
	"os"
	"time"
)

func Num9() {
	in := make(chan int)
	out := make(chan int)
	i := 0
	go Reader(in, out)

	for {
		in <- i
		i++
		res := <-out
		fmt.Fprintln(os.Stdout, res)
		time.Sleep(time.Second)

	}

}

func Reader(ch1, ch2 chan int) {
	for {
		num := <-ch1
		ch2 <- num * 2
	}

}
