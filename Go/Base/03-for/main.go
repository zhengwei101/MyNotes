package main

import (
	"fmt"
)

func count() int {
	fmt.Println("count:")
	return 3
}

func getData() []int {
	fmt.Println("origin data.")
	return []int{10, 20, 30}
}

// 在go里面，没有while循环，只有唯一的for循环
func main() {
	//初始化语句仅被执行一次。
	for i, c := 0, count(); i < c; i++ { //初始化语句的count函数仅执行一次
		fmt.Println("a", i)
	}
	c := 0
	for c < count() {
		println("b", c) //条件表达式中的count重复执行
		c++
	}
	//最简单的for循环就是在for后面什么都不写，代表一个死循环
	for {
		fmt.Println("loop")
		break // 循环途中可以用break跳出
	}

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue // 用continue继续循环
		}
		fmt.Println("odd:", n)
	}

	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//在迭代遍历时，for...range 返回索引、键值数据
	//for...range支持字符串、数组、数组指针、切片、字典、通道类型
	println("***for...range***")
	x := []int{100, 200, 300}
	for i, n := range x {
		println(i, ":", n)
	}

	//允许返回单值，或用"_"忽略
	data := [3]string{"a", "b", "c"}
	for i := range data { //只返回1st value
		fmt.Println(i, data[i])
	}

	for _, s := range data { //忽略1st value
		fmt.Println(s)
	}

	for range data { //仅迭代，不返回。可用来执行清空channel等操作

	}

	for i, s := range data {
		fmt.Println(&i, &s)
	}

	{
		data := [3]int{10, 20, 30}
		for i, x := range data {
			if i == 0 {
				data[0] += 100
				data[1] += 200
				data[2] += 300
			}
			fmt.Printf("x: %d,data: %d\n", x, data[i])
		}

		for i, x := range data[:] {
			if i == 0 {
				data[0] += 100
				data[1] += 200
				data[2] += 300
			}
			fmt.Printf("x: %d,data: %d\n", x, data[i])
		}

		//输出
		// x: 10,data: 110
		// x: 20,data: 220
		// x: 30,data: 330
		// x: 110,data: 210 //当i==0修改data时，x已经取值，所以是110
		// x: 420,data: 420
		// x: 630,data: 630
	}

	//range目标表达式也可以是函数
	{
		for i, x := range getData() {
			fmt.Println(i, x)
		}
	}

	//配合标签，break和continue可在多层嵌套中指定目标层级
	println("***for...break...continue***")
	{
	outer:
		for x := 0; i < 5; x++ {
			for y := 0; y < 10; y++ {
				if y > 2 {
					continue outer
				}
				if x > 2 {
					break outer
				}
				fmt.Print(x, ":", y, " ")
			}
		}
	}
}
