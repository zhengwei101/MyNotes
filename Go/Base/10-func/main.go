package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func add2(a, b int) int {
	return a + b
}

func exists(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}

func test(x int) func() { //返回函数类型
	return func() { //匿名函数
		println(x) //闭包
	}
}

//return语句不是ret汇编指令，它会先更新返回值
func testDefer() (z int) {
	defer func() {
		println("defer: ", z)
		z += 100 //修改命名返回值
	}()

	return 100 //实际执行次序：z=100, call defer, ret
}

// 变参本质上就是一个切片，只能接收一到多个同类型参数，且必须放在列表尾部
func testArgs(s string, a ...int) {
	fmt.Printf("%T, %v, %v\n", a, a, s)
}

func testArgs2(a ...int) {
	for i := range a {
		a[i] += 100
	}
}

func testFuncParam(s string, f func(string)) {
	f(s)
}

func testFuncReturn() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func testStruct() {
	type calc struct { //定义结构体类型
		mul func(x, y int) int //函数类型字段
	}

	x := calc{
		mul: func(x, y int) int {
			return x * y
		},
	}
	fmt.Printf("struct: %d\n", x.mul(2, 3))
}

func testChannel() {
	c := make(chan func(int, int) int, 2)

	c <- func(x, y int) int {
		return x + y
	}
	fmt.Printf("channel: %d\n", (<-c)(2, 2))
}

func testClosure() []func() {
	var s []func()

	for i := 0; i < 2; i++ {
		s = append(s, func() { //将多个匿名函数添加到列表
			println(&i, i)
		})
	}
	return s //返回匿名函数列表
}

//多个匿名函数引用同一环境变量，会让事情变得更加复杂
//任何的修改行为都会影响其他函数取值，在并发模式下可能需要做同步处理
func testMulti(x int) (func(), func()) { //返回两个匿名函数
	return func() {
			println(x)
			x += 10 //修改环境变量
		}, func() {
			println(x) //显示环境变量
		}
}

func main() {
	res := add(1, 2)
	fmt.Println(res, add2(100, 200))

	//匿名函数
	{
		//直接执行：
		func(s string) {
			fmt.Println(s)
		}("hello,world")

		//赋值给变量
		add := func(x, y int) int {
			return x + y
		}
		fmt.Printf("add: %d\n", add(11, 22))

		//作为返回值
		add2 := testFuncReturn()
		fmt.Printf("add2: %d\n", add2(11, 22))

		//作为参数
		testFuncParam("hello,world again", func(s string) {
			fmt.Println(s)
		})
	}

	testStruct()
	testChannel()

	//迭代执行所有匿名函数
	for _, f := range testClosure() {
		f()
	}

	f1, f2 := testMulti(100)
	f1()
	f2()

	//函数原生支持返回多个值
	v, ok := exists(map[string]string{"a": "A"}, "a")
	fmt.Println(v, ok) // A True

	x := 100
	f := test(x)
	f()

	println("testDefer:", testDefer())

	testArgs("abc", 1, 2, 3)

	{
		//将切片作为变参时，须进行展开操作。如果是数组，先将其转换为切片
		a := [3]int{10, 20, 30}
		testArgs("def", a[:]...) //转换为slice后展开

		//既然变参是切片，那么参数复制的仅是切片自身，并不包括底层数组，因此可以修改原数据
		//如果需要，可用内置函数copy复制㡳层数据
		testArgs2(a[:]...)
		fmt.Println(a)
	}
}
