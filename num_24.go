package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func Dist(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64((p1.x-p2.x)), 2) + math.Pow(float64(p1.y-p2.y), 2))
}

func Num24() {
	p1 := Point{4, 5}
	p2 := Point{1, 1}
	fmt.Println(Dist(p1, p2))
}
