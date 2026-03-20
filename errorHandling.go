package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Sentinel error used with errors.Is.
var ErrInvalidAge = errors.New("invalid age")

// ValidationError is a custom error type used with errors.As.
type ValidationError struct {
	Field string
	Msg   string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("validation failed for %s: %s", v.Field, v.Msg)
}

func checkNumber(num int) (string, error) {
	if num < 0 {
		return "", fmt.Errorf("number must be >= 0, got %d", num)
	}
	return "number is valid", nil
}

func parseAge(input string) (int, error) {
	age, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("parse age %q: %w", input, err)
	}
	if age <= 0 {
		return 0, fmt.Errorf("%w: age must be greater than zero", ErrInvalidAge)
	}
	if age > 130 {
		return 0, ValidationError{Field: "age", Msg: "must be <= 130"}
	}
	return age, nil
}

func readConfig(path string) error {
	_, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read config %s: %w", path, err)
	}
	return nil
}

func main() {
	fmt.Println("1) Basic error return")
	result, err := checkNumber(-5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println("\n2) Error wrapping and unwrapping")
	_, err = parseAge("abc")
	if err != nil {
		fmt.Println("Wrapped Error:", err)
		fmt.Printf("Unwrapped error: %v\n", errors.Unwrap(err))
	}

	fmt.Println("\n3) errors.Is with sentinel error")
	_, err = parseAge("0")
	if err != nil {
		if errors.Is(err, ErrInvalidAge) {
			fmt.Println("Matched sentinel error ErrInvalidAge:", err)
		} else {
			fmt.Println("Different error:", err)
		}
	}

	fmt.Println("\n4) errors.As with custom error type")
	_, err = parseAge("200")
	if err != nil {
		var ve ValidationError
		if errors.As(err, &ve) {
			fmt.Printf("Validation Error -> field: %s, msg: %s\n", ve.Field, ve.Msg)
		} else {
			fmt.Println("Non-validation error:", err)
		}
	}

	fmt.Println("\n5) Checking file-not-found")
	err = readConfig("missing-config.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Config file does not exist:", err)
		} else {
			fmt.Println("Read error:", err)
		}
	}
}
