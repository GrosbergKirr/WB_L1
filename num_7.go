package main

import (
	"context"
	"fmt"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func Num7() {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	var workerCount int
	fmt.Println("Enter workers amount: ")
	_, err := fmt.Scanln(&workerCount)
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}
	ch := make(chan int)
	mu := sync.Mutex{}
	dict := map[int]string{}

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
					mu.Lock()
					dict[res] = strconv.Itoa(res)
					fmt.Printf("hello from gourotine %d: take number %d\n", i, res)
					mu.Unlock()
				}
			}
		}()

	}

	r := 0
	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			fmt.Println(dict)
			return
		default:
			ch <- r
			r++
			time.Sleep(time.Second)
		}

	}

}
