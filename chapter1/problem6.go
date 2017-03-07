package chapter1

import "strconv"

// StringCompression compresses string
// Time complexity O(n)
// Space complexity O(1)
func StringCompression(str string) string {
	if len(str) == 0 {
		return str
	}

	bs := make([]byte, 0)
	pre := str[0]
	count := 1
	for i := 1; i < len(str); i++ {
		if str[i] != pre {
			bs = append(bs, pre)
			bs = append(bs, []byte(strconv.Itoa(count))...)
			pre = str[i]
			count = 1
		} else {
			count++
		}
	}
	bs = append(bs, pre)
	bs = append(bs, []byte(strconv.Itoa(count))...)

	if len(str) < len(bs) {
		return str
	}
	return string(bs)
}
