package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) == 0 {
		return ""
	}

	// Count chars needed
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	
	have := make(map[byte]int)
	required := len(need)
	formed := 0

	l := 0
	minLen := math.MaxInt32
	start := 0

	for r := 0; r < len(s); r++ {
		ch := s[r]
		have[ch]++

		// If char count matches what's needed
		if count, ok := need[ch]; ok && have[ch] == count {
			formed++
		}

		// Try shrinking window
		for formed == required {
			// Update smallest window
			if r-l+1 < minLen {
				minLen = r - l + 1
				start = l
			}

			// Pop from left
			leftChar := s[l]
			have[leftChar]--
			if count, ok := need[leftChar]; ok && have[leftChar] < count {
				formed--
			}
			l++
		}
	}

	if minLen == math.MaxInt32 {
		return ""
	}
	return s[start : start+minLen]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t)) // Output: "BANC"
}
