package quiz_test

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// 数组是可以比较的
func TestArrayComparson(t *testing.T) {
	a := [...]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	t.Log(a == b) // true
}

// 切片是不可以比较的，运行后报错
// func TestSliceComparson(t *testing.T) {
// 	a := []int{1, 2, 3}
// 	b := []int{1, 2, 3}
// 	t.Log(a == b) //slice can only be compared to ni
// }

type Something struct {
	i int
	s string
}

// Go中结构体是可以比较的，比较成员的值
func TestComparson(t *testing.T) {
	c1 := Something{1, "1"}
	d1 := Something{1, "1"}
	c2 := &Something{1, "1"}
	d2 := &Something{1, "1"}
	t.Log(c1 == d1, c2 == d2)
}

// Go 语⾔是传值的，所以数组会被复制，修改不会影响原有数组。
func TestRefArray(t *testing.T) {
	a := [3]int{1, 2, 3}
	fn := func(x [3]int) { x[1] = 4 }
	fn(a)
	t.Log(a) //[1,2,3]
}

// 切片是结构体，虽然Go是传值的，
// 但是复制后的结构体和原结构体都包含指向切片元素的存储空间的指针,
// 并且由于复制，他们指向的是同一空间
func TestRefSlice(t *testing.T) {
	a := []int{1, 2, 3}
	fn := func(x []int) { x[1] = 4 }
	fn(a)
	t.Log(a) //[1,4,3]
}

func TestForRange(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}
	//var dst []*int
	dst := []*int{}

	for _, i := range src {
		dst = append(dst, &i)
	}
	for _, p := range dst {
		fmt.Print(*p)
	}
}

func TestPanicVxExit(t *testing.T) {
	fmt.Println("Start")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()

	os.Exit(1)
	fmt.Println("End")
}

func runTask(id int) string {
	time.Sleep(10 * time.Microsecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string) //存在协和泄露，应改为 ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func GetFn() func() {
	fmt.Println("[outside]")
	return func() {
		fmt.Println("[inside]")
	}
}

func TestDefer(t *testing.T) {
	defer GetFn()()
	fmt.Print("[here]")
}

// 3 2 1 0
func TestDefer2(t *testing.T) {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

// 函数选项模式
// 初始化一个类的时候，如果类的构造参数较多，尽量使用Option写法
type Foo struct {
	name string
	id   int
	age  int
	db   interface{}
}

// FooOption代表可选参数,将foo指针传递进去，能让内部函数进行修改
type FooOption func(foo *Foo)

// WithName 代表Name为可选参数
func WithName(name string) FooOption {
	return func(foo *Foo) {
		foo.name = name
	}
}

// WithAge 代表Age为可选参数
func WithAge(age int) FooOption {
	return func(foo *Foo) {
		foo.age = age
	}
}

// WithDB 代表db为可选参数
func WithDB(db interface{}) FooOption {
	return func(foo *Foo) {
		foo.db = db
	}
}

// NewFoo 代表初始化，只有一个必填字段id
func NewFoo(id int, options ...FooOption) *Foo {
	//按照默认值初始化一个foo对象
	foo := &Foo{
		name: "default",
		id:   id,
		age:  10,
		db:   nil,
	}
	//遍历options改造这个foo对象
	for _, option := range options {
		option(foo)
	}
	return foo
}

// 具体使用NewFoo的函数
func TestBar(t *testing.T) {
	foo := NewFoo(1, WithAge(15))
	fmt.Println(foo)
}
