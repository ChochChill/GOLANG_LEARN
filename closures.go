package main

import "fmt"

// Basic closure — captures variable from outer scope
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Closure with parameter — adder factory
func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// Closure capturing a mutable slice
func makeAccumulator() func(int) []int {
	collected := []int{}
	return func(n int) []int {
		collected = append(collected, n)
		return collected
	}
}

// Closure used as a one-time transformer with state
func makeMultiplierWithHistory(factor int) (func(int) int, func() []int) {
	history := []int{}
	multiply := func(n int) int {
		result := n * factor
		history = append(history, result)
		return result
	}
	getHistory := func() []int {
		return history
	}
	return multiply, getHistory
}

// Closure fixing the classic loop-capture bug
func makeGreetersCorrected(names []string) []func() string {
	greeters := make([]func() string, len(names))
	for i, name := range names {
		name := name // capture loop variable correctly
		greeters[i] = func() string {
			return "Hello, " + name
		}
	}
	return greeters
}

// Closure as middleware — wraps a function with logging
func withLogging(label string, fn func(int) int) func(int) int {
	return func(n int) int {
		result := fn(n)
		fmt.Printf("[%s] input=%d output=%d\n", label, n, result)
		return result
	}
}

func main() {
	// Independent counters — each has its own captured state
	counterA := makeCounter()
	counterB := makeCounter()

	fmt.Println("Counter A:", counterA()) // 1
	fmt.Println("Counter A:", counterA()) // 2
	fmt.Println("Counter A:", counterA()) // 3
	fmt.Println("Counter B:", counterB()) // 1 (independent)

	// Adder closure
	add5 := makeAdder(5)
	add10 := makeAdder(10)

	fmt.Println("add5(3):", add5(3))   // 8
	fmt.Println("add10(3):", add10(3)) // 13

	// Accumulator closure — captures and grows a slice
	accumulate := makeAccumulator()
	fmt.Println(accumulate(10)) // [10]
	fmt.Println(accumulate(20)) // [10 20]
	fmt.Println(accumulate(30)) // [10 20 30]

	// Closure pair sharing the same captured state
	double, getHistory := makeMultiplierWithHistory(2)
	fmt.Println("double(4):", double(4))  // 8
	fmt.Println("double(7):", double(7))  // 14
	fmt.Println("History:", getHistory()) // [8 14]

	// Correct loop-capture with closures
	greeters := makeGreetersCorrected([]string{"Alice", "Bob", "Charlie"})
	for _, greet := range greeters {
		fmt.Println(greet()) // Hello, Alice / Bob / Charlie
	}

	// Closure wrapping another function with logging
	square := func(n int) int { return n * n }
	loggedSquare := withLogging("square", square)
	loggedSquare(5) // [square] input=5 output=25
	loggedSquare(9) // [square] input=9 output=81
}
