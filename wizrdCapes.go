package main

import "fmt"

func main() {
	var count int
	fmt.Scan(&count)
	for count > 0 {
		var length int
		fmt.Scan(&length)
		arr := make([]int, length)
		for i:=0; i<length; i++{
			fmt.Scan(&arr[i])
		}
		var combinations = length
		for i:0; i<length; i++{
			tempC := length - arr[i] + 1
			if(combinations>tempC){
				combinations=tempC	
			}
		}
		fmt.Println(combinations)
		count--
	}
}
