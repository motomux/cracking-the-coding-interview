package main

import "testing"

func TestCheckPermutation(t *testing.T) {
	tests := map[string]struct {
		str1, str2 string
		out        bool
	}{

		"should return true when given strings are the same": {
			str1: "hello",
			str2: "hello",
			out:  true,
		},

		"should return true when given strings are permutation": {
			str1: "hello",
			str2: "leloh",
			out:  true,
		},

		"should return false when given strings are not permutation": {
			str1: "hello",
			str2: "oeloh",
			out:  false,
		},

		"should return false when given strings are length": {
			str1: "hello",
			str2: "helloo",
			out:  false,
		},

		"should return true when given strings contains unicode": {
			str1: "こんにちは、世界",
			str2: "世界、こんにちは",
			out:  true,
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := CheckPermutation(test.str1, test.str2)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}
