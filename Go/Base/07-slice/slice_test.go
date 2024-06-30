package main_test

import (
	"fmt"
	"testing"
)

type TransInfo interface{}

type Fragment interface {
	Exec(transInfo TransInfo) error
}
type GetPodAction struct{}

func (g GetPodAction) Exec(transInfo TransInfo) error {
	return nil
}

func TestSlice(t *testing.T) {
	//s := make([]int)
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := make([]int, 5, 10)
	s3 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1, s2, s3)
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i) //追加数据。当超过容量限制时，自动分配更大的存储空间
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	// if a == b { //切片只能和nil比较
	// 	t.Log("equal")
	// }
	t.Log(a, b)
}

// func TestSliceStack(t *testing.T) {
// 	//栈最大容量5
// 	stack := make([]int, 0, 5)
// }

func TestSliceExtention(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	t.Log(s1) //[2 3 4 5]
	t.Log(s2) //[5 6]
	//slice是数组arr的视图(view)
	//slice可以向后扩展，不可以向前扩展
	//s[i]不可以超越len(s), 向后扩展不可以超载㡳层数组cap(s)
	s3 := append(s2, 10) //将7改为10
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3,s4,s5 = ", s3, s4, s5)
	fmt.Println("arr = ", arr)
	//添加元素时，如果超越cap，系统会重新分配更大的底层数组
}

func SliceAppend(s []int) []int {
	s = append(s, 3)
	return s
}

func TestSliceAppend(t *testing.T) {
	s := make([]int, 0)
	fmt.Println(s)
	s = SliceAppend(s)
	fmt.Println(s)

	course := []string{"go", "grpc", "mysql", "es", "gin"}
	//mycourse := append(course[:2], course[3:]...) //这里会修改course的数据
	mycourse := make([]string, 1) //这样则不会修改course的数据
	mycourse = append(mycourse, course[:2]...)
	mycourse = append(mycourse, course[3:]...)
	fmt.Println("mycourse:", mycourse)
	fmt.Println("course: ", course)
	course = course[:3]
	fmt.Println("reduce: ", course)

}
