package main

import "fmt"

func bubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

func insertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && a[j-1] > a[j]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	//取最左边的值为基准值
	val := arr[left]
	// a[l+1...k] < v; arr[k+1...i] > v
	k := left
	// partition并确定val所在的位置
	for i := left + 1; i <= right; i++ {
		if arr[i] < val {
			arr[k+1], arr[i] = arr[i], arr[k+1]
			k++
		}
	}
	// swap value
	arr[left], arr[k] = arr[k], arr[left]
	quickSort(arr, left, k-1)
	quickSort(arr, k+1, right)
}

func selectionSort(a []int) {

}
func main() {
	b := [...]int{8, 7, 5, 4, 3, 10, 15}
	bubbleSort(b[:])
	fmt.Println("bubble sort:", b)

	b = [...]int{8, 7, 5, 4, 3, 10, 15}
	insertionSort(b[:])
	fmt.Println("insertion sort:", b)

	b = [...]int{8, 7, 5, 4, 3, 10, 15}
	quickSort(b[:], 0, len(b)-1)
	fmt.Println("quick sort:", b)
}
