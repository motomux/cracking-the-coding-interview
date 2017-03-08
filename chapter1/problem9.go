package chapter1

import "strings"

// StringRotation checks if str2 is a rotation of str1
func StringRotation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	return strings.Contains(str1+str1, str2)
}
