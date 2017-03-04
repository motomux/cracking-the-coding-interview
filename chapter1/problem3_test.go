package chapter1

import "testing"

func TestURLify(t *testing.T) {
	tests := map[string]struct {
		str string
		pos int
		out string
	}{

		"should return empty string when empty string is given": {
			str: "",
			pos: 0,
			out: "",
		},

		"should return encoded string when given string includes a space": {
			str: "hello world  ",
			pos: 10,
			out: "hello%20world",
		},

		"should return the same string when given string doesn't include a space": {
			str: "helloworld",
			pos: 9,
			out: "helloworld",
		},

		"should return encoded string when given string includes multiple spaces": {
			str: "hello world !!    ",
			pos: 13,
			out: "hello%20world%20!!",
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			out := URLify(test.str, test.pos)
			if out != test.out {
				t.Errorf("actual=%q expected=%q", out, test.out)
			}
		})
	}
}
