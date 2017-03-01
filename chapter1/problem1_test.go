package main

import "testing"

func TestIsUnique(t *testing.T) {
	tests := map[string]struct {
		in  string
		out bool
	}{
		"should return true if string contains only unique characters": {
			in:  "world",
			out: true,
		},

		"should return false if string contains duplicated characters": {
			in:  "hello world",
			out: false,
		},

		"should return true if empty string is given": {
			in:  "",
			out: true,
		},

		"should return true if string contains only unique unicode characters": {
			in:  "こんにちは、世界",
			out: true,
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := IsUnique(test.in)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}
