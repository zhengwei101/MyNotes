package interface_test

import (
	"fmt"
	"testing"
)

// 接口采用了duck type方式，也就是说无须在实现类型上添加显示声明
type user struct {
	name string
	age  byte
}

func (u user) Print() {
	fmt.Printf("%+v\n", u)
}

type Printer interface { //接口类型
	Print()
}

type Programmer interface { //接口类型
	WriteHello() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHello() string {
	return "fmt.Println(\"Hello world\")"
}

func TestClient(t *testing.T) {
	var u user
	u.name = "Tom"
	u.age = 35

	var x Printer = u //只要包含接口所需的全部方法，即表示实现了该接口
	x.Print()

	var p Programmer = new(GoProgrammer)
	t.Log(p.WriteHello())

	//另有空接口类型interface{}, 用途类似OOP里的system.Object,可接收任意类型对象。
}
