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

//用defer定义延迟调用，无论函数是否出错，它都确保结束前被调用
func testDefer(a, b int) {
	defer println("dispose A...") // 常用来释放资源，解除锁定，或执行一些清理操作
	defer println("dispose B...") //可定义多个defer，按FILO顺序执行
	println(a / b)
}

// go里面的函数原生支持返回多个值
func main() {
	res := add(1, 2)
	fmt.Println(res) // 3

	v, ok := exists(map[string]string{"a": "A"}, "a")
	fmt.Println(v, ok) // A True

	x := 100
	f := test(x)
	f()

	//testDefer(10, 0) //panic: runtime error: integer divide by zero
}
