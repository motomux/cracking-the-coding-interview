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
