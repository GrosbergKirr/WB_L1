package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Num6() {

	var workerTime int
	fmt.Println("Enter work time: ")
	_, err := fmt.Scanln(&workerTime)
	if err != nil {
		panic(err)
	}

	chCanc := make(chan int)
	ctxCanc, cancel := context.WithCancel(context.Background())
	ctxTime, _ := context.WithTimeout(context.Background(), time.Second*5)
	wg := sync.WaitGroup{}
	count := 0

	wg.Add(1)
	go ChanCancel(chCanc, &wg, count)

	wg.Add(1)
	go CtxCancel(ctxCanc, &wg, count)

	wg.Add(1)
	go CtxTimeOut(ctxTime, &wg, count)

	time.Sleep(time.Duration(workerTime) * time.Second)

	chCanc <- -1

	cancel()

	wg.Done()

}

func CtxTimeOut(ctx context.Context, wg *sync.WaitGroup, count int) {
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			fmt.Println("Time goroutine stop")
			return
		default:
			fmt.Printf("time goroutine work: %d\n", count)
			count++
			time.Sleep(time.Second)
		}
	}
}
func CtxCancel(ctx context.Context, wg *sync.WaitGroup, count int) {
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			fmt.Println("Cancel goroutine stop")
			return
		default:
			fmt.Printf("cancel goroutine work: %d\n", count)
			count++
			time.Sleep(time.Second)
		}
	}
}

func ChanCancel(ch chan int, wg *sync.WaitGroup, count int) {
	for {
		select {
		case <-ch:
			wg.Done()
			fmt.Println("Chan goroutine stop")
			return
		default:
			fmt.Printf("chan goroutine work: %d\n", count)
			count++
			time.Sleep(time.Second)
		}
	}
}
