package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")

	slice1 := make([]int, 6)
	slice2 := make([]int, 6, 10)
	fmt.Println(slice1, slice2)
	fmt.Println(len(slice1), len(slice2))
	fmt.Println(cap(slice1), cap(slice2))

	array := []int{6, 8, 10}
	var ret = minValue(array)
	fmt.Printf("最小值是： %d\n", ret)

	{
	OuterLoop:
		for i := 0; i < 2; i++ {
			for j := 0; j < 5; j++ {
				switch j {
				case 3:
					fmt.Println(i, j)
					continue OuterLoop
				}
			}
		}
	}
	A()
	B()
	C()

	//通过goto语句来实现标签跳转
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto BreakTag
			}
			fmt.Println(i, " ", j)
		}
	}
BreakTag:
	fmt.Println("Done")
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

func minValue(arr []int) (m int) {
	m = arr[0]
	for _, v := range arr {
		if v < m {
			m = v
		}
	}
	return
}
