package main

import (
	"fmt"
	"net/http"
)

func Hello(a ...int) string {
	return fmt.Sprintf("Hello, %v!", a)
}

func add_numbers(a int, b int) int {
	return a + b
}

func GetUrl(url string) (resp string, err error) {
	resphttp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resphttp.Body.Close()
	if resphttp.StatusCode != http.StatusOK {
		err = fmt.Errorf("failed to fetch URL: %s, status code: %d", url, resphttp.StatusCode)
		return
	}
	resp = fmt.Sprintf("Status Code: %d", resphttp.StatusCode)
	return
}
func main() {
	fmt.Println(Hello(1, 2, 3))

	url := "https://www.google.com"
	result, err := GetUrl(url)
	if err != nil {
		fmt.Printf("Error fetching URL: %v\n", err)
	} else {
		fmt.Println(result)
	}

	url = "https://www.nonexistentwebsite.com"
	result, err = GetUrl(url)
	if err != nil {
		fmt.Printf("Error fetching URL: %v\n", err)
	} else {
	
		fmt.Println(result)
	}
	
}
