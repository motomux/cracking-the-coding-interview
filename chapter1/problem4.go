package chapter1

// CheckPalindromePermutation checks if given string is palindrome permutation
// Time complexity O(n)
// Space complexity O(n)
func CheckPalindromePermutation(str string) bool {
	counts := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		counts[str[i]]++
	}

	oddCount := 0
	for _, count := range counts {
		if count%2 == 1 {
			oddCount++
		}
	}

	if oddCount > 1 {
		return false
	}
	return true
}

// CheckPalindromePermutation2 checks if given string is palindrome permutation
// It can check only string with lowers
// Time complexity O(n)
// Space complexity O(1)
func CheckPalindromePermutation2(str string) bool {
	table := 0

	for i := 0; i < len(str); i++ {
		table ^= 1 << uint(str[i]-'a')
	}

	if (table-1)&table == 0 {
		return true
	}
	return false
}
