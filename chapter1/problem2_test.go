package chapter1

import "testing"

func TestCheckPermutation(t *testing.T) {
	tests := []struct {
		str1, str2 string
		out        bool
	}{
		{
			str1: "hello",
			str2: "hello",
			out:  true,
		},
		{
			str1: "hello",
			str2: "leloh",
			out:  true,
		},
		{
			str1: "hello",
			str2: "oeloh",
			out:  false,
		},
		{
			str1: "hello",
			str2: "helloo",
			out:  false,
		},
		{
			str1: "こんにちは、世界",
			str2: "世界、こんにちは",
			out:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.str1+"_"+test.str2, func(t *testing.T) {
			out := CheckPermutation(test.str1, test.str2)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}

func TestCheckPermutation2(t *testing.T) {
	tests := []struct {
		str1, str2 string
		out        bool
	}{
		{
			str1: "hello",
			str2: "hello",
			out:  true,
		},
		{
			str1: "hello",
			str2: "leloh",
			out:  true,
		},
		{
			str1: "hello",
			str2: "oeloh",
			out:  false,
		},
		{
			str1: "hello",
			str2: "helloo",
			out:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.str1+"_"+test.str2, func(t *testing.T) {
			out := CheckPermutation2(test.str1, test.str2)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}
