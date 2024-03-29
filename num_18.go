package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	count int64
}

func Num18() {
	c := new(Counter)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go Incrementor(c, &wg)
		//time.Sleep(time.Second)
	}
	wg.Wait()
	fmt.Println(c.count)
}

func Incrementor(c *Counter, wg *sync.WaitGroup) {
	atomic.AddInt64(&c.count, 1)
	wg.Done()

}
