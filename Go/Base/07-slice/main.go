package main

import "fmt"

//slice的本质是一个数组片段的描述
//包括数组指针
//片段的长度
//片段的容量

//切片不同于数组，可以任意更改长度，然后也有更多丰富的操作
//可以用make来创建一个切片，可以像数组一样去取值，使用append来追加元素
//注意append的用法，你必须把append的结果赋值回原slice。
//在你执行append的时候，如果容量不够的话，它会扩容并返回新的slice
//slice初始化时，也可以指定长度
func main() {
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get:", s[2])   // c
	fmt.Println("len:", len(s)) // 3

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s) // [a b c d e f]

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c) // [a b c d e f]

	fmt.Println(s[2:5]) // [c d e]
	fmt.Println(s[:5])  // [a b c d e]
	fmt.Println(s[2:])  // [c d e f]

	good := []string{"g", "o", "o", "d"}
	fmt.Println(good) // [g o o d]
}
