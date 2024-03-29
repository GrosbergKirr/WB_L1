package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Num4() {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	var workerCount int
	fmt.Println("Enter workers amount: ")
	_, err := fmt.Scanln(&workerCount)
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}
	ch := make(chan int)

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("gourotine %d stoped\n", i)
					wg.Done()
					return
				case res := <-ch:
					fmt.Printf("hello from gourotine %d: take number %d\n", i, res)
				}
			}
		}()

	}

	i := 0
	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			return
		default:
			ch <- i
			i++
			time.Sleep(time.Second)
		}

	}

}
