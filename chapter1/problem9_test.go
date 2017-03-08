package chapter1

import "testing"

func TestStringRotation(t *testing.T) {
	tests := map[string]struct {
		str1, str2 string
		out        bool
	}{

		"should return true when given inputs are empty": {
			str1: "",
			str2: "",
			out:  true,
		},

		"should return true when str2 is rotation of str1": {
			str1: "abcde",
			str2: "cdeab",
			out:  true,
		},

		"should return false when str2 is not rotation of str1": {
			str1: "abcde",
			str2: "cdefa",
			out:  false,
		},

		"should return false when str2 is rotation of str1 but they don't have the same length": {
			str1: "abcde",
			str2: "cdea",
			out:  false,
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := StringRotation(test.str1, test.str2)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}
