package function

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"advance/function/fib"
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

// 使用闭包输出斐波那契数列
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

func TestFibonacci(t *testing.T) {
	f := fib.Fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func writeFile(filename string) {
	//file, err := os.Create(filename)
	file, err := os.OpenFile(
		filename,
		os.O_EXCL|os.O_CREATE,
		0666,
	)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			//如果err不是PathError，就结束程序
			panic(err)
		} else {
			fmt.Println("PathError:", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func TestWriteFile(t *testing.T) {
	writeFile("fib.txt")
}

func fibonacci2() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

// 为函数实现Read接口
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 100 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	//TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func TestFibonacci2(t *testing.T) {
	f := fibonacci2()
	printFileContents(f)
}
