package main

import (
	"errors"
	"fmt"
	"strconv"
)

type user struct {
	name     string
	password string
}

// 在函数的返回值类型里面，加一个error，就代表这个函数可能会返回错误
// 在函数实现的时候，return需要返回两个值
// 出现错误，就return nil和error
// 没有错误，就return 结果和nil
func findUser(users []user, name string) (v *user, err error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}
func NewError(text string) error {
	return &errorString{text} //这里要取地地址
}

var ErrType = NewError("EOF")

func main() {
	if ErrType == NewError("EOF") {
		fmt.Println("Error: ", ErrType)
	}

	x, y := 1, 2
	x, y = y+3, x+2 //先计算出右值y+3、x+2，然后再对x,y变量赋值
	fmt.Println(x, y)

	//空标识符_(blank identifier)，通过作为忽略占位符使用
	i, _ := strconv.Atoi("12") //忽略Atoi的error返回值
	fmt.Println(i)

	//在处理函数错误时，退化赋值允许我们重复使用err变量，这是相当有益的
	w, err := findUser([]user{{"wang", "1024"}}, "wang")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(w.name) // wang

	//err退化为赋值，l为新定义
	if l, err := findUser([]user{{"li", "2048"}}, "wang"); err != nil {
		fmt.Println(err) // not found
		return
	} else {
		fmt.Println(l.name)
	}
}
