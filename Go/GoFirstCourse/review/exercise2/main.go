package main

import (
	"fmt"
	"strconv"
)

//值类型： 基本数据类型int、float、bool、string以及数组和struct
//引用类型： 指针、slice、map、chan等都是引用类型

func Add(n1, n2, n3 int) (int, int) {
	result := n1 + n2 + n3
	quotient := result / 10
	remainder := result % 10
	return remainder, quotient
}

// 大数相加
func BigNumAdd(num1 string, num2 string) string {
	//num1 890243589034859254954989214935904389809852
	//num2                    98984253247384728753248
	//num3                   153973468183289118563100
	//num3 890243589034859255053973468183289118563100

	//num1  95498921493590438980985
	//num2   9898425324738472875324
	//num3 105397346818328911856309

	len1 := len(num1)
	len2 := len(num2)
	var len int = 0

	if len1 >= len2 {
		len = len1
	} else {
		len = len2
	}

	if len <= 0 {
		return "0"
	}

	var result string
	var thirdNum int = 0

	for i := 0; i < len; i++ {
		var (
			n1, n2 int
			err    error
			r1     int
		)

		if i < len1 {
			n1, err = strconv.Atoi(string(num1[len1-i-1]))
			if err != nil {
				break
			}
		}

		if i < len2 {
			n2, err = strconv.Atoi(string(num2[len2-i-1]))
			if err != nil {
				break
			}
		}

		r1, thirdNum = Add(n1, n2, thirdNum)
		result = fmt.Sprintf("%d%s", r1, result)
		//fmt.Println(result)
	}

	if thirdNum != 0 {
		result = fmt.Sprintf("%d%s", thirdNum, result)
	}

	return result
}

// 九九乘法表
func multi() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}

// 判断一段字符中是否为回文字符串
func process(str string) bool {
	r := []rune(str)
	length := len(r)
	for i := range r {
		if i == length/2 {
			break
		}
		last := length - i - 1
		if r[i] != r[last] {
			return false
		}
	}
	return true
}

func callProcess() {
	var str string
	fmt.Scanf("%sd", &str)
	if process(str) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

// 字符统计
func counter(str string) (wordCount, spaceCount, numberCount, otherCount int) {
	r := []rune(str)
	for _, v := range r {
		switch {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			wordCount++
		case v == ' ':
			spaceCount++
		case v >= '0' && v <= '9':
			numberCount++
		default:
			otherCount++
		}
	}
	return
}

func callCounter() {
	var str string = "abc 上海自来水 123"
	wc, sc, nc, oc := counter(str)
	fmt.Printf("word count: %d\t,space count: %d\t,number count: %d\t,other count: %d\t", wc, sc, nc, oc)
}

func main() {
	// str := "abbacba"
	// fmt.Println(strings.Trim(str, "ab")) //c

	//fmt.Println(BigNumAdd("890243589034859254954989214935904389809852", "98984253247384728753248"))
	fmt.Println(BigNumAdd("95498921493590438980985", "9898425324738472875324"))

	//multi()

	//callProcess()

	//callCounter()
}
