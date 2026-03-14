package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
    s1Len := len(s1)
    if s1Len > len(s2) {
        return false
    }

    counts := make([]int, 26)
    for i, c := range s1 {
        counts[c - 'a']++
        counts[s2[i] - 'a']--
    }

    var matches int
    for _, count := range counts {
        if count == 0 {
            matches++
        }
    }

    left, right := 0, len(s1)
    for right < len(s2) {
        if matches == 26 {
            return true
        }

        leftKey := s2[left]-'a'
        counts[leftKey]++
        if counts[leftKey] - 1 == 0 {
            matches--
        } else if counts[leftKey] == 0 {
            matches++
        }
        left++

        rightKey := s2[right]-'a'
        counts[rightKey]--
        if counts[rightKey] + 1 == 0 {
            matches--
        } else if counts[rightKey] == 0 {
            matches++
        }
        right++
    }

    return matches == 26
}

func main() {
	tests := []struct {
		s1, s2 string
		want   bool
	}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
		{"hello", "ooolleoooleh", false},
		{"trinitrophenylmethylnitramine", "dinitrophenylhydrazinetrinitrophenylmethylnitramine", true},
	}

	for _, tt := range tests {
		got := checkInclusion(tt.s1, tt.s2)
		fmt.Printf("s1=%q, s2=%q -> %v (want %v)\n", tt.s1, tt.s2, got, tt.want)
	}
}
 