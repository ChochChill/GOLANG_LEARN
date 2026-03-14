package main

import("fmt")

func main()  {
	arr := []int{1,2,1,3,4,5}
	fmt.Println(hasDuplicate((arr)))
	arr = []int{1,2,0,3,4,5}
	fmt.Println(hasDuplicate((arr)))
}

func hasDuplicate(arr []int) bool{
	count := make(map[int]bool)
	for _, n := range arr{
		ok := count[n]
		if(ok){
			return true
		}
		count[n]=true
	}
	return false 
}