package main

import (
	"fmt"
	"os"
)

// 1. Resource Cleanup (The most common use case)
func cleanupExample() {
	fmt.Println("\n--- Resource Cleanup ---")
	file, err := os.Create("temp.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	// Always close the file, even if something fails later in the function
	defer func() {
		fmt.Println("Closing the file now...")
		file.Close()
		os.Remove("temp.txt") // Clean up the temp file
	}()

	fmt.Println("File created and being used...")
}

// 2. LIFO (Last-In-First-Out) Order
func lifoExample() {
	fmt.Println("\n--- LIFO Order ---")
	fmt.Println("Counting starts:")
	
	for i := 1; i <= 3; i++ {
		// Defers are pushed onto a stack
		defer fmt.Printf("Deferred call %d\n", i)
	}
	
	fmt.Println("Function logic finished, now running defers...")
}

// 3. Parameter Evaluation Timing
func evaluationExample() {
	fmt.Println("\n--- Parameter Evaluation ---")
	i := 0
	
	// The value of 'i' is captured WHEN the defer is called, not when it runs
	defer fmt.Println("Value of 'i' in deferred call (captured early):", i)
	
	i = 10
	fmt.Println("Current value of 'i':", i)
}

// 4. Modifying Named Return Values
func namedReturnExample() (result int) {
	fmt.Println("\n--- Named Return Modification ---")
	defer func() {
		// Defers can access and modify named return variables after 'return'
		result += 5
		fmt.Println("Defer added 5 to the result")
	}()
	
	return 10 // Returns 10 + 5 = 15
}

// 5. Recovering from a Panic
func recoverExample() {
	fmt.Println("\n--- Recover from Panic ---")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from a scary panic:", r)
		}
	}()

	fmt.Println("About to do something dangerous...")
	panic("BOOM!") // This would normally crash the program
	fmt.Println("This line will never run")
}

func main() {
	fmt.Println("Starting Defer Use Case Examples")

	cleanupExample()
	lifoExample()
	evaluationExample()
	
	val := namedReturnExample()
	fmt.Println("Final returned value:", val)

	recoverExample()

	fmt.Println("\nProgram finished safely!")
}
