package main

import "fmt"

type Human struct {
	Name    string
	Surname string
	Age     int8
}

type Action struct {
	Human
}

func (h *Human) Greet() {
	fmt.Printf("Hi! My name is %s %s", h.Name, h.Surname)
}

func (h *Human) HowOld() {
	fmt.Printf("I'm %d years old\n", h.Age)
}

func Num1() {
	h := Human{"Ivan", "Ivanov", 25}
	h.Greet()
	h.HowOld()

	a := Action{h}
	a.Greet()
}
