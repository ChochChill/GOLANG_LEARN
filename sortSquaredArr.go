package main

import "fmt"

func main() {
	var count int
	fmt.Scan(&count)
	for count > 0 {
		var length int
		fmt.Scan(&length)
		arr := make([]int, length)
		for i := 0; i < length; i++ {
			fmt.Scan(&arr[i])
		}
		fmt.Print(sortedSquares(arr, length))
		count--
	}
}

func sortedSquares(nums []int, length int) []int {
	i, j := 0, length-1
	for i < j {
		eleL := nums[i] * nums[i]
		eleR := nums[j] * nums[j]
		if eleL > eleR {
			nums[j] = nums[i]
			nums[i] = eleR
			i++
		} else {
			nums[j] = eleR
			j--
		}
	}
	nums[i] = nums[i] * nums[i]
	return nums
}
