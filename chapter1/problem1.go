package main

// IsUnique checks given string contains only unique characters
// Time complexity O(n)
// Space complesity O(n)
func IsUnique(str string) bool {
	table := make(map[rune]bool, 0)

	for _, r := range str {
		if _, ok := table[r]; ok {
			return false
		}
		table[r] = true
	}

	return true
}

// IsUnique2 checks given string contains only unique characters
// It can check only lowercase of alphabet, otherwise result can be wrong
// Time complexity O(n)
// Space complesity O(1)
func IsUnique2(str string) bool {
	var table int

	for i := 0; i < len(str); i++ {
		c := str[i]
		if c < 'a' && 'z' < c {
			return false
		} else if table&(1<<uint(c-'a')) > 1 {
			return false
		}
		table = table | (1 << uint(c-'a'))
	}

	return true
}
