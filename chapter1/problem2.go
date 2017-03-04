package chapter1

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

// CheckPermutation2 checks given strings are permutation
// It can check only lowercase of alphabet, otherwise result can be wrong
// Time complexity O(n)
// Space complexity O(1)
func CheckPermutation2(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	table := 0
	for i := 0; i < len(str1); i++ {
		c := str1[i]
		if c < 'a' && 'z' < c {
			return false
		}
		table ^= 1 << uint(c-'a')
	}

	for i := 0; i < len(str2); i++ {
		c := str2[i]
		if c < 'a' && 'z' < c {
			return false
		}
		table ^= 1 << uint(c-'a')
	}

	if table == 0 {
		return true
	}
	return false
}
