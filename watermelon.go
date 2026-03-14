package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	if number > 2 && number%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
