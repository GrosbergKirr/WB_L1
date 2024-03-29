package main

import (
	"fmt"
	"time"
)

func Num25() {
	var sleeptime int64
	fmt.Println("Enter workers amount: ")
	_, err := fmt.Scanln(&sleeptime)
	if err != nil {
		panic(err)
	}
	MySleep(sleeptime)
}
func MySleep(a int64) {
	t := time.Now().Unix()
	fmt.Println(t)
	for time.Now().Unix() < t+a {

	}
	fmt.Println(time.Now().Unix() - t)
}
