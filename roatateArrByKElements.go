package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 4, 5, 6, 7, 0}
	k := 3
	reverse(arr, 0, k-1)
	reverse(arr, k, len(arr)-1)
	reverse(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func reverse(arr []int, beg int, end int) {
	for end > beg {
		arr[beg], arr[end] = arr[end], arr[beg]
		beg++
		end--
	}
}
