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

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func TestQueue(t *testing.T) {
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	q.Push("abc")
	fmt.Println(q.Pop())
}

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func TestFuncClosure(t *testing.T) {
	a := adder()
	for i := 1; i < 10; i++ {
		fmt.Printf("0+1+...+%d=%d\n", i, a(i))
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func TestFuncClosure2(t *testing.T) {
	a := adder2(0)
	for i := 1; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0+1+...+%d=%d\n", i, s)
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func TestFibonacci(t *testing.T) {
	f := fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}
