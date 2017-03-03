package main

// CheckPermutation checks given strings are permutation
// Time complexity O(n)
// Space complexity O(n)
func CheckPermutation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	table := make(map[rune]int, 0)
	for _, r := range str1 {
		if _, ok := table[r]; !ok {
			table[r] = 0
		}
		table[r]++
	}

	for _, r := range str2 {
		if _, ok := table[r]; !ok {
			return false
		}
		table[r]--
	}

	for _, v := range table {
		if v != 0 {
			return false
		}
	}

	return true
}
