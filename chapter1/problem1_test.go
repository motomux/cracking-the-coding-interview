package chapter1

import "testing"

func TestIsUnique(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{
			in:  "world",
			out: true,
		},
		{
			in:  "hello world",
			out: false,
		},
		{
			in:  "",
			out: true,
		},
		{
			in:  "こんにちは、世界",
			out: true,
		},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			out := IsUnique(test.in)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}

func TestIsUnique2(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{
			in:  "world",
			out: true,
		},
		{
			in:  "helloworld",
			out: false,
		},
		{
			in:  "",
			out: true,
		},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			out := IsUnique2(test.in)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}
