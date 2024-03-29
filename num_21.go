package main

import (
	"fmt"
)

type OpticFiber struct{}

func (dog *OpticFiber) Opto() {
	fmt.Print("Photons")
}

type GPONSocket struct {
	OpticFiber
}

func (adapter *GPONSocket) OptToLan() {
	fmt.Print("Get: ")
	adapter.Opto()
	fmt.Println("And convert it to electrons (casual LAN)")
}

func Num21() {
	d := GPONSocket{}
	d.OptToLan()
}
