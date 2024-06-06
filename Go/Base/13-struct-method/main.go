package main

import "fmt"

//可以为当前包类的任意类型定义方法
type X int

func (x *X) increase() {
	*x++
}

type user struct {
	name     string
	password string
}

func (u user) ToString() string {
	return fmt.Sprintf("%+v", u)
}
func (u user) checkPassword(password string) bool {
	return u.password == password
}

func (u *user) resetPassword(password string) {
	u.password = password
}

type manager struct {
	user
	title string
}

// go里面可以为结构体定义一些方法
// 具体的代码修改，就是把第一个参数，加上括号，写到函数名称前面。
// 在实现结构体的方法的时候也有两种写法，一种是带指针，一种不带指针。
// 带指针的话，就可以对这个结构体的变量做修改
// 不带指针的话，实际上操作的是一个拷贝，无法对变量做修改。
func main() {
	var x X
	x.increase()
	println(x)

	var m manager
	m.name = "Tom"
	m.password = "123"
	m.title = "CEO"

	println(m.ToString()) //调用user.ToString()

	a := user{name: "wang", password: "1024"}
	a.resetPassword("2048")
	fmt.Println(a.checkPassword("2048")) // true
}
