package main

import "fmt"

// Higher-order function
func operate(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// Functions to pass
func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

// Higher-order function that returns a function
func multiplier(factor int) func(int) int {
	return func(value int) int {
		return factor * value
	}
}

func process(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = fn(v)
	}
	return result
}

func square(n int) int {
	return n * n
}

// Factory function to create greeting functions
func greeting(language string) func(string) string {
	switch language {
	case "English":
		return func(name string) string {
			return "Hello, " + name
		}
	case "Spanish":
		return func(name string) string {
			return "Hola, " + name
		}
	default:
		return func(name string) string {
			return "Hi, " + name
		}
	}
}

// Filter function
func filter(nums []int, predicate func(int) bool) []int {
	result := []int{}
	for _, v := range nums {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Predicate functions
func isEven(n int) bool {
	return n%2 == 0
}

func isGreaterThanThree(n int) bool {
	return n > 3
}

func main() {
	// Using the higher-order function
	fmt.Println("Sum:", operate(3, 5, add))          // Output: Sum: 8
	fmt.Println("Product:", operate(3, 5, multiply)) // Output: Product: 15

	double := multiplier(2) // Returns a function to double the input
	triple := multiplier(3) // Returns a function to triple the input

	fmt.Println("Double of 5:", double(5)) // Output: Double of 5: 10
	fmt.Println("Triple of 5:", triple(5)) // Output: Triple of 5: 15

	numbers := []int{1, 2, 3, 4, 5}
	squared := process(numbers, square)
	fmt.Println("Squared Numbers:", squared) // Output: [1 4 9 16 25]

	greetInEnglish := greeting("English")
	greetInSpanish := greeting("Spanish")

	fmt.Println(greetInEnglish("Alice")) // Output: Hello, Alice
	fmt.Println(greetInSpanish("Bob"))   // Output: Hola, Bob

	numbers = []int{1, 2, 3, 4, 5, 6}
	evens := filter(numbers, isEven)
	greaterThanThree := filter(numbers, isGreaterThanThree)

	fmt.Println("Even Numbers:", evens)           // Output: [2 4 6]
	fmt.Println("Numbers > 3:", greaterThanThree) // Output: [4 5 6]
}
