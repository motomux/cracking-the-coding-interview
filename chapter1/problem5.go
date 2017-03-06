package chapter1

// OneAway checks if given strings are one away
// Time complexity O(n)
// Space complexity O(1)
func OneAway(str1, str2 string) bool {
	if diff := len(str1) - len(str2); diff == 0 {
		return checkReplace(str1, str2)
	} else if diff == 1 {
		return checkAdd(str1, str2)
	} else if diff == -1 {
		return checkAdd(str2, str1)
	}
	return false
}

func checkReplace(str1, str2 string) bool {
	diff := 0

	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			diff++
		}
	}

	if diff == 1 {
		return true
	}
	return false
}

func checkAdd(str1, str2 string) bool {
	diff := 0

	for i, k := 0, 0; i < len(str1) && k < len(str2); {
		if str1[i] != str2[k] {
			diff++
			i++
		} else {
			i++
			k++
		}
	}

	if diff == 1 || diff == 0 {
		return true
	}
	return false
}
