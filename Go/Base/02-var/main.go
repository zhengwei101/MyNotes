package main

import (
	"fmt"
	"math"
)

// 字符串、整数、布尔型、浮点型等
func main() {
	var a = "initial"
	var b, c int = 1, 2
	var d = true
	var e float64
	f := float32(e)
	g := a + "foo"
	fmt.Println(a, b, c, d, e, f) // initial 1 2 true 0 0
	fmt.Println(g)                // initialfoo

	const s string = "constant"
	const h = 500000000
	const i = 3e20 / h
	//constant 500000000 6e+11 -0.28470407323754404 0.7591864109375384
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
}
