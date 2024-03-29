package main

import (
	"fmt"
	"math/big"
)

func Num22() {
	a := new(big.Int)
	a.SetString("9000000000000000000000000000000", 10)
	b := new(big.Int)
	b.SetString("3000000000000000000000000000000", 10)
	sm := new(big.Int)
	df := new(big.Int)
	ml := new(big.Int)
	dv := new(big.Int)
	sm.Add(a, b)
	df.Sub(a, b)
	ml.Mul(a, b)
	dv.Div(a, b)
	fmt.Println(sm, df, ml, dv)
}
