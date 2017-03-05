package chapter1

import "testing"

func TestCheckPalindromePermutation(t *testing.T) {
	tests := map[string]struct {
		in  string
		out bool
	}{
		"should return true when empty string is given": {
			in:  "",
			out: true,
		},

		"should return true when palindrome permutation is given": {
			in:  "tteeiiii",
			out: true,
		},

		"should return true when palindrome permutation is given with odd length": {
			in:  "tteeiiiis",
			out: true,
		},

		"should return false when non palindrome permutation is given": {
			in:  "tteeiiab",
			out: false,
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := CheckPalindromePermutation(test.in)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}

func TestCheckPalindromePermutation2(t *testing.T) {
	tests := map[string]struct {
		in  string
		out bool
	}{
		"should return true when empty string is given": {
			in:  "",
			out: true,
		},

		"should return true when palindrome permutation is given": {
			in:  "tteeiiii",
			out: true,
		},

		"should return true when palindrome permutation is given with odd length": {
			in:  "tteeiiiis",
			out: true,
		},

		"should return false when non palindrome permutation is given": {
			in:  "tteeiiab",
			out: false,
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := CheckPalindromePermutation2(test.in)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}
