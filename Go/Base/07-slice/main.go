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
	fmt.Println("get:", s[2], "len:", len(s)) // c 3

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

	s1 := make([]int, 3, 5)           //指定len、cap，底层数组初始化为0
	s2 := make([]int, 3)              //省略cap，和len相等
	s3 := []int{10, 20, 5: 30}        //按初始化元素分配㡳层数组，并设置len、cap
	fmt.Println(s1, len(s1), cap(s1)) //[0 0 0] 3 5
	fmt.Println(s2, len(s2), cap(s2)) //[0 0 0] 3 3
	fmt.Println(s3, len(s3), cap(s3)) //[10 20 0 0 0 30] 6 6

	//注意下面这两种定义方式的区别，前者仅定义了一个[]int类型变量，并未执行初始化操作
	//而后者则用初始化表达式完成了全部创建过程
	var a []int
	b := []int{}
	println(a == nil, b == nil)

	s4 := []int{0, 1, 2, 3, 4}
	p := &s4     //取header地址
	p0 := &s4[0] //取array[0]
	p1 := &s4[1]
	println(p, p0, p1)
	(*p)[0] += 100 //*[]int不支持索引操作，须先返回[]int对象
	*p1 += 100     //直接用元素指针操作
	fmt.Println(s4)
}
