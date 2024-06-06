package main

import (
	"fmt"
	"math"
	"strconv"
)

// 自定义类型
type flags byte

const (
	read flags = 1 << iota
	write
	exec
)

// 字符串、整数、布尔型、浮点型等
func main() {
	var a = "initial"
	var b, c int = 1, 2
	var d = true
	var e float64
	f := float32(e)
	g := a + "foo"
	fmt.Println(a, b, c, d, e, f) // initial 1 2 true 0 0
	fmt.Println(g)                // initialfoo

	const s string = "constant"
	const h = 500000000
	const i = 3e20 / h
	//constant 500000000 6e+11 -0.28470407323754404 0.7591864109375384
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))

	{
		//二元操作符
		//除位移操作外，操作数两边类型必须相同。
		//如果其中一个是无显示类型声明的常量，那么该常量操作数会自动转型
		const v = 20 //无显示类型声明的常量
		var a byte = 10
		b := v + a //v自动转换为byte/uint8类型
		fmt.Printf("auto convert: %T,%v\n", b, b)

		const c float32 = 1.2
		d := c + v //v自动转换为float32类型
		fmt.Printf("auto convert: %T,%v\n", d, d)
	}
	{
		//支持八进制、十六进制以及科学记数法
		a, b, c := 100, 0144, 0x64
		fmt.Println(a, b, c)
		fmt.Printf("0b%b,%#o,%#x\n", a, a, a)
		fmt.Println(math.MinInt8, math.MaxInt8)
	}

	{
		//AND NOT 按位清除(bit clear) a &^ b , 0110 &^ 1011 = 0100
		const (
			read byte = 1 << iota
			write
			exec
			freeze
		)
		a := read | write | freeze
		b := read | freeze | exec
		c := a &^ b
		fmt.Printf("bit clear: %04b &^ %04b = %04b\n", a, b, c)
	}

	//标准库strconv可在不同进制（字符串）间转换
	{
		a, _ := strconv.ParseInt("1100100", 2, 32)
		b, _ := strconv.ParseInt("0144", 8, 32)
		c, _ := strconv.ParseInt("64", 16, 32)
		println(a, b, c)
		fmt.Println("0b" + strconv.FormatInt(a, 2))
		fmt.Println("0" + strconv.FormatInt(b, 8))
		fmt.Println("0x" + strconv.FormatInt(c, 16))
	}

	{
		flag := read | exec
		fmt.Printf("%b\n", flag) // 输出二进制标记位

		//和var、const类似，多个type定义可合并成组，可在函数或代码块内定义局部类型
		type ( //一组自定义类型
			user struct { //结构体类型
				name string
				age  int
			}
			event func(string) bool //函数类型
		)

		u := user{"Tom", 22}
		fmt.Println(u)

		var f event = func(s string) bool {
			fmt.Println(s)
			return s != ""
		}
		f("abc")
	}

}
