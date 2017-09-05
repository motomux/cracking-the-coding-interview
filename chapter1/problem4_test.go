package chapter1

import "testing"

func TestCheckPalindromePermutation(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{
			in:  "",
			out: true,
		},
		{
			in:  "tteeiiii",
			out: true,
		},
		{
			in:  "tteeiiiis",
			out: true,
		},
		{
			in:  "tteeiiab",
			out: false,
		},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			out := CheckPalindromePermutation(test.in)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}

func TestCheckPalindromePermutation2(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{
			in:  "",
			out: true,
		},
		{
			in:  "tteeiiii",
			out: true,
		},
		{
			in:  "tteeiiiis",
			out: true,
		},
		{
			in:  "tteeiiab",
			out: false,
		},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			out := CheckPalindromePermutation2(test.in)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}
