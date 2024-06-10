package main

import "fmt"

func main() {
	var a [5]int              //元素自动初始化为0
	b := [4]int{2, 5}         //未提供初始值的元素自动初始化为0
	c := [4]int{5, 3: 10}     //可指定索引位置初始化
	d := [...]int{1, 2, 3}    //编译器按初始化值数量确定数组长度
	e := [...]int{10, 3: 100} //支持索引初始化，注意数组长度与此有关
	fmt.Println(a, b, c, d, e)
	a[4] = 100
	fmt.Println("get:", a[2])
	fmt.Println("len:", len(a))

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
