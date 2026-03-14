package main

import "fmt"

var (
	s      string
	prefix [26][]int // prefix[c][i] = count of ('a'+c) in s[:i]
	n      int
)

// getCount counts occurrences of character ch in s[l:r)
func getCount(ch byte, l, r int) int {
	return prefix[ch-'a'][r] - prefix[ch-'a'][l]
}

// Precompute prefix sums for all characters
func buildPrefix() {
	for c := 0; c < 26; c++ {
		prefix[c] = make([]int, n+1)
	}
	for i := 0; i < n; i++ {
		for c := 0; c < 26; c++ {
			prefix[c][i+1] = prefix[c][i]
		}
		prefix[s[i]-'a'][i+1]++
	}
}

// Recursive function to count minimal changes
func countMoves(ch byte, beg, length int) int {
	if length == 1 {
		if s[beg] == ch {
			return 0
		}
		return 1
	}

	mid := beg + length/2
	leftCount := getCount(ch, beg, mid)
	rightCount := getCount(ch, mid, beg+length)

	costIfLeftGood := (length/2 - leftCount) + countMoves(ch+1, mid, length/2)
	costIfRightGood := (length/2 - rightCount) + countMoves(ch+1, beg, length/2)

	if costIfLeftGood < costIfRightGood {
		return costIfLeftGood
	}
	return costIfRightGood
}

func main() {
	var count int
	fmt.Scan(&count)
	for count > 0 {
		var length int
		fmt.Scan(&length)
		fmt.Scan(&s)
		n = len(s)
		buildPrefix()
		fmt.Println(countMoves('a', 0, n))
		count--
	}

}
