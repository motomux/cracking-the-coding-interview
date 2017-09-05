package chapter1

import "testing"

func TestStringCompression(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			in:  "",
			out: "",
		},
		{
			in:  "abbcccdddeeeee",
			out: "a1b2c3d3e5",
		},
		{
			in:  "abcde",
			out: "abcde",
		},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			out := StringCompression(test.in)
			if out != test.out {
				t.Errorf("actual=%s expected=%s", out, test.out)
			}
		})
	}
}
