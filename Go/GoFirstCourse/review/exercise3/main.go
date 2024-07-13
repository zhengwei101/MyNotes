package main

import (
	"fmt"
	"strings"
)

// 闭包使用示例
func makeSuffix(suffix string) func(string) string {
	return func(fileName string) string {
		if !strings.HasSuffix(fileName, suffix) {
			return fileName + suffix
		}
		return fileName
	}
}

func main() {
	f1 := makeSuffix(".bmp")
	f2 := makeSuffix(".png")

	fmt.Println(f1("apple"))
	fmt.Println(f2("banana"))
}
