package main

import (
	"fmt"
)

func main() {
	nums:= []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(getMaxSubArray(nums))
}

func getMaxSubArray(nums []int) (int, []int){
	if(len(nums)==0){
		return 0,[]int{}
	}
	maxSum:= nums[0]
	curSum:= nums[0]
	l,r,tempL:=0,0,0
	for i:=1; i<len(nums); i++{
		n := nums[i]
		if(curSum+n<n){
			curSum=n
			tempL=i
		}	else{
			curSum=curSum+n
		}
		if(curSum>maxSum){
			maxSum=curSum
			l=tempL
			r=i
		}
	}
	return maxSum, nums[l:r+1]
}