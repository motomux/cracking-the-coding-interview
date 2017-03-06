package chapter1

import "testing"

func TestOneAway(t *testing.T) {
	tests := map[string]struct {
		str1, str2 string
		out        bool
	}{

		"should return false when given strings are empty": {
			str1: "",
			str2: "",
			out:  false,
		},

		"should return false when given strings are the same": {
			str1: "hello",
			str2: "hello",
			out:  false,
		},

		"should return true when given strings has the same length and one chars which are different": {
			str1: "hello",
			str2: "hella",
			out:  true,
		},

		"should return false when given strings has the same length and two chars which are different": {
			str1: "hello",
			str2: "helaa",
			out:  false,
		},

		"should return true when str1 is one char longer than str2 and has the same string with one additional": {
			str1: "hello",
			str2: "hell",
			out:  true,
		},

		"should return true when str1 is one char longer than str2 and has one away string": {
			str1: "heell",
			str2: "hell",
			out:  true,
		},

		"should return true when str2 is one char longer than str1 and has the same string with one additional": {
			str1: "hell",
			str2: "hello",
			out:  true,
		},

		"should return true when str2 is one char longer than str1 and has one away string": {
			str1: "hell",
			str2: "heell",
			out:  true,
		},

		"should return false when difference b/w length of given strings is more than 1": {
			str1: "hello",
			str2: "hellooo",
			out:  false,
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := OneAway(test.str1, test.str2)
			if out != test.out {
				t.Errorf("actual=%t expected=%t", out, test.out)
			}
		})
	}
}
