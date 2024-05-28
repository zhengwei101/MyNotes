package main

import "fmt"

// 在go里面，没有while循环，只有唯一的for循环,
//最简单的for循环就是在for后面什么都不写，代表一个死循环
//循环途中可以用break跳出,用continue继续循环
func main() {

	i := 1
	for {
		fmt.Println("loop")
		break
	}
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
}
