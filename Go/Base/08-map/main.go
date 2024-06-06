package main

import "fmt"

func main() {
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
