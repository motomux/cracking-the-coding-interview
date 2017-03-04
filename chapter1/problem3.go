package chapter1

// URLify replace space with %20
// Time complexity O(n)
// Space complexity O(0)
func URLify(str string, pos int) string {
	res := []byte(str)
	for i := len(res) - 1; i >= 0 && pos >= 0; i, pos = i-1, pos-1 {
		if res[pos] == ' ' {
			res[i] = '0'
			res[i-1] = '2'
			res[i-2] = '%'
			i -= 2
		} else {
			res[i] = res[pos]
		}
	}

	return string(res)
}
