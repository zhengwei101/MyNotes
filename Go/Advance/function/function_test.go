package function

import (
	"fmt"
	"testing"
	"time"
)

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent: ", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestFunction(t *testing.T) {
	tsSF := timeSpent(slowFun)
	fmt.Printf("result: %d\n", tsSF(10))
	fmt.Printf("sum: %d", Sum(1, 2, 3, 4, 5))
}
