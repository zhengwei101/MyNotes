package main

import (
	"fmt"
	"net/http"
	"slices"
	"sync"
)

func ConCat() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	arr3 := []int{7, 8, 9}
	merged := append(arr1, arr2...) //v1.22之前合并切片
	merged = append(merged, arr3...)
	fmt.Println(merged)

	merged = slices.Concat(arr1, arr2, arr3) //v1.22更方便
	fmt.Println(merged)
}

func NewHttpMethod() {
	//http://127.0.0.1:5678/hello/ddd
	mux := http.NewServeMux()
	//restful风格参数
	mux.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "hello "+r.PathValue("name"))
		if err != nil {
			return
		}
	})
}

func main() {
	wg := sync.WaitGroup{}
	values := []string{"a", "b", "c"}
	wg.Add(len(values))
	for _, v := range values { //每次for range 都会为v重新申请一块内存
		fmt.Printf("%p\n", &v)
		go func() {
			defer wg.Done()
			fmt.Println(v)
		}()
	}
	wg.Wait()

	for i := range 3 { //v1.22 新的遍历方式
		fmt.Println(i)
	}

	ConCat()
}
