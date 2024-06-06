package defer_test

import (
	"testing"
)

func TestDefer(t *testing.T) {
	//用defer定义延迟调用，无论函数是否出错，它都确保结束前被调用
	defer func() {
		t.Log("Clean resources") // 常用来释放资源，解除锁定，或执行一些清理操作
	}()

	x, y := 1, 2
	defer func(a int) { //可定义多个defer，按FILO顺序执行
		t.Log("defer x,y= ", a, y) //y为闭包引用
	}(x) //注册时复制调用参数

	x += 100 //对x的修改不会影响延迟调用
	y += 200
	t.Log(x, y)

	t.Log("Started")
	panic("Fatal error") //defer 仍会执行
}
