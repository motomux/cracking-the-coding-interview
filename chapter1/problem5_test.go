package chapter1

import "testing"

func TestOneAway(t *testing.T) {
	tests := []struct {
		str1, str2 string
		out        bool
	}{
		{
			str1: "",
			str2: "",
			out:  false,
		},
		{
			str1: "hello",
			str2: "hello",
			out:  false,
		},
		{
			str1: "hello",
			str2: "hella",
			out:  true,
		},
		{
			str1: "hello",
			str2: "helaa",
			out:  false,
		},
		{
			str1: "hello",
			str2: "hell",
			out:  true,
		},
		{
			str1: "heell",
			str2: "hell",
			out:  true,
		},
		{
			str1: "hell",
			str2: "hello",
			out:  true,
		},
		{
			str1: "hell",
			str2: "heell",
			out:  true,
		},
		{
			str1: "hello",
			str2: "hellooo",
			out:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.str1+"_"+test.str2, func(t *testing.T) {
			out := OneAway(test.str1, test.str2)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}
