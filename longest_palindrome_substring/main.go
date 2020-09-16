package main

// task description
// https://leetcode.com/problems/longest-palindromic-substring/submissions/

/* Expand Around Center */
func getPalindrome(s []rune, start, end int) (pal string) {
	for (start >= 0 && end < len(s)) && (s[start] == s[end]) {
		pal = string(s[start : end+1])
		start--
		end++
	}
	return pal
}

func longestPalindrome_ExpandAroundCenter(s string) string {
	if len(s) <= 1 {
		return s
	}

	longest := ""
	for i := range s {
		pal := getPalindrome([]rune(s), i, i)
		if len(longest) <= len(pal) {
			longest = pal
		}
		pal = getPalindrome([]rune(s), i, i+1)
		if len(longest) <= len(pal) {
			longest = pal
		}
	}
	return longest
}

/* Dynamic programming */
func longestPalindrome_DynamicProgramming(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	// init
	table := make([][]bool, n)
	for i := 0; i < n; i++ {
		table[i] = make([]bool, n)
	}

	// filling the main diagonal
	for i := 0; i < n; i++ {
		table[i][i] = true
	}
	longest := string(s[0])

	// filling the table for palindromes of length = 2
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			table[i][i+1] = true
			longest = s[i : i+1+1]
		}
	}

	// filling the table for palindromes of length greater than 2
	for k := 3; k <= n; k++ {
		for i := 0; i < n-k+1; i++ {
			j := i + k - 1
			if table[i+1][j-1] && s[i] == s[j] {
				table[i][j] = true

				if k >= len(longest) {
					longest = s[i : j+1]
				}
			}
		}
	}
	return longest
}
