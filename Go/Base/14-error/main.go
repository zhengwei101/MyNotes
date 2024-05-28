package main

import (
	"errors"
	"fmt"
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

func main() {
	u, err := findUser([]user{{"wang", "1024"}}, "wang")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name) // wang

	if u, err := findUser([]user{{"wang", "1024"}}, "li"); err != nil {
		fmt.Println(err) // not found
		return
	} else {
		fmt.Println(u.name)
	}
}
