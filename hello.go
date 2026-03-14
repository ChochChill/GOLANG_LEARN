package main

import (
	"fmt"
)

func main() {
	var a int = 10
	b := 20 // short declaration
	const pi float64 = 3.14

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("pi:", pi)

	//for loop on intergers in go
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	i := 0
	//while loop in go
	for i < 5 {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
		i++
	}

	//arr: fixed length
	//slice: dynamic length
	var arr [3]int = [3]int{1, 2, 3} // array
	slice := []int{4, 5, 6}          // slice
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)
	slice = append(slice, 7)
	fmt.Println("after append Slice:", slice)
	// Slice another slice
	sub := slice[1:3]
	fmt.Println("Sub-slice:", sub)
	// Length and capacity
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))
	//capacity is initial length slice can hold before it needs to allocate more memory.
	data := make([]int, 0, 1000) // allocates enough space for 1000 elements from the start, length given is 0
	fmt.Println("custom slice length:", len(data), "& capacity:", cap(data))

	//map (dict)
	m := make(map[string]int)
	// Insert
	m["apple"] = 10
	m["banana"] = 20
	// Access
	fmt.Println("Apple:", m["apple"])
	// Check existence
	val, ok := m["grape"]
	fmt.Println("Grape exists?", ok, "Value:", val)
	// Delete
	delete(m, "banana")
	fmt.Println("Map after deletion:", m)
	q := make([]int, 3) // initialized slice: [0, 0, 0]
	fmt.Println(q)
	p := new([]int) // pointer to nil slice
	fmt.Println(p)
}
