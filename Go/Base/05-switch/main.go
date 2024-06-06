package main

import (
	"fmt"
	"time"
)

// go语言里的switch不用加break
// 可以在switch后面不加任何变量，然后在case里写条件分支
func main() {
	x := 2
	switch x { //将x与case条件做匹配
	case 1: //单个匹配条件
		fmt.Println("one")
	case 2, 3: //多个匹配条件，命中其一即可(OR)
		fmt.Println("two or three")
	default:
		fmt.Println("other")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	{
		//switch同样支持初始化语句，按从上到下、从左到右顺序匹配case执行。
		//只有全部匹配失败，才会执行default块
		switch x := 5; x {
		default: //编译器确保不会先执行default块
			x += 100
			fmt.Println(x)
		case 5:
			x += 50
			fmt.Println(x)
		}
	}
}
