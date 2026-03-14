package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 1, 0, 3, 12}
	pos := 0
	for _, n := range arr {
		if n != 0 {
			arr[pos] = n
			pos++
		}
	}

	for pos < len(arr) {
		arr[pos] = 0
		pos++
	}

	fmt.Println(arr)
}
