package chapter1

import "testing"

func TestStringRotation(t *testing.T) {
	tests := []struct {
		str1, str2 string
		out        bool
	}{
		{
			str1: "",
			str2: "",
			out:  true,
		},
		{
			str1: "abcde",
			str2: "cdeab",
			out:  true,
		},
		{
			str1: "abcde",
			str2: "cdefa",
			out:  false,
		},
		{
			str1: "abcde",
			str2: "cdea",
			out:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.str1+"_"+test.str2, func(t *testing.T) {
			out := StringRotation(test.str1, test.str2)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}
