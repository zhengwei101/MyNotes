package main

import (
	"fmt"
	"math/rand"
)

func MyRand() {
	for i := 0; i < 3; i++ {
		a := rand.Intn(100)
		fmt.Println(a)
	}
}

func Add(a int, b int, c chan int) {
	c <- (a + b)
}

func CallAdd() {
	c := make(chan int, 1)
	go Add(2, 5, c)
	sum := <-c
	fmt.Println("sum=", sum)
}

func MyReverse(str string) string {
	var result string
	strlen := len(str)
	for i := 0; i < strlen; i++ {
		result += fmt.Sprintf("%c", str[strlen-i-1])
	}
	return result
}
func MyReverse1(str string) string {
	var result []byte
	strBytes := []byte(str)
	len := len(strBytes)
	for i := 0; i < len; i++ {
		result = append(result, strBytes[len-i-1])
	}
	return string(result)
}

func main() {
	MyRand()
	str := MyReverse("hello world")
	fmt.Println(str)
	str = MyReverse1(str)
	fmt.Println(str)

	CallAdd()
}
