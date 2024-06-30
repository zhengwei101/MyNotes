package main

import "fmt"

func mapFunc[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a), cap(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}

func filterFunc[T any](a []T, f func(T) bool) []T {
	var n []T
	for _, e := range a {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}

func testGeneric() {
	vi := []int{1, 2, 3, 4, 5, 6}
	vs := mapFunc(vi, func(v int) string {
		return "<" + fmt.Sprint(v) + ">"
	})
	fmt.Println(vs)

	vf := filterFunc(vi, func(v int) bool {
		return v < 4
	})
	fmt.Println(vf)
}

func main() {
	testGeneric()

	m := make(map[string]int) //创建字典类型对象
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)           // map[one:1 two:2]
	fmt.Println(len(m))      // 2
	fmt.Println(m["one"])    // 1
	fmt.Println(m["unknow"]) // 0

	for k, v := range m {
		fmt.Printf("%s, %d\n", k, v)
	}

	// 使用ok-idiom获取值，可知道key/value是否存在
	if r, ok := m["unknow"]; !ok {
		fmt.Println("key 'unknow' is not existing")
		fmt.Println(r, ok) // 0 false
	}

	//所谓ok-idiom，是指在多返回值中用一个名为ok的布尔值来标示操作是否成功。

	delete(m, "one") //删除
	fmt.Println(m)   // map[two:2]
	m2 := map[string]int{"one": 1, "two": 2}
	var m3 = map[string]int{"one": 1, "two": 2}
	fmt.Println(m2, m3) // map[one:1 two:2] map[one:1 two:2]
}
