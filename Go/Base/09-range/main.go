package main

import "fmt"

// 对于slice或map,我们可以用range来快速遍历。
// range遍历的时候，对于数组会返回两个值，一个是索引，一个是对应位置的值。
// 不需要索引的话，可能用下划来忽略。
func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		if num == 2 {
			fmt.Println("index:", i, "num:", num) // index: 0 num: 2
		}
	}
	fmt.Println(sum) // 9

	m := map[string]string{"a": "A", "b": "B"}
	for k, v := range m {
		fmt.Println(k, v) // b 8; a A
	}
	for k := range m {
		fmt.Println("key", k) // key a; key b
	}
}
