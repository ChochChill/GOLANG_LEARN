package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 3, 2, 2, 1, 5, 1, 6, 1}
	k := 5
	fmt.Println(kMostFreqElements(arr, k))
}

func kMostFreqElements(arr []int, k int) []int {
	count := make(map[int]int)
	for _, n := range arr {
		_, ok := count[n]
		if ok {
			count[n]++
		} else {
			count[n] = 1
		}
	}
	buckets := make([][]int, len(arr)+1)
	for k, v := range count {
		buckets[v] = append(buckets[v], k)
	}
	result := []int{}
	for i := len(buckets) - 1; i >= 0; i-- {
		for _, ele := range buckets[i] {
			result = append(result, ele)
			if len(result) == k {
				return result
			}
		}
	}
	return result
}
