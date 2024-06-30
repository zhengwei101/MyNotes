package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")

	A()
	B()
	C()
}

func A() {
	fmt.Println("Func A")
}

func B() {
	fmt.Println("Func B")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()

	panic("Panic in B")
}

func C() {
	fmt.Println("Func C")
}
