package main

import "fmt"

// 在go里面，没有while循环，只有唯一的for循环
func main() {
	//最简单的for循环就是在for后面什么都不写，代表一个死循环
	for {
		fmt.Println("loop")
		break // 循环途中可以用break跳出
	}

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue // 用continue继续循环
		}
		fmt.Println(n)
	}

	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//在迭代遍历时， for...range 除元素外，还可返回索引
	x := []int{100, 200, 300}
	for i, n := range x {
		println(i, ":", n)
	}
}
