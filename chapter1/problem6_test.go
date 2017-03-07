package chapter1

import "testing"

func TestStringCompression(t *testing.T) {
	tests := map[string]struct {
		in  string
		out string
	}{

		"should return empty string when empty string is given": {
			in:  "",
			out: "",
		},

		"should return compressed string": {
			in:  "abbcccdddeeeee",
			out: "a1b2c3d3e5",
		},

		"should return original string if compressed string is longer than original": {
			in:  "abcde",
			out: "abcde",
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := StringCompression(test.in)
			if out != test.out {
				t.Errorf("actual=%s expected=%s", out, test.out)
			}
		})
	}
}
