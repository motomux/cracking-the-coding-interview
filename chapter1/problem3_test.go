package chapter1

import (
	"strconv"
	"testing"
)

func TestURLify(t *testing.T) {
	tests := []struct {
		str string
		pos int
		out string
	}{
		{
			str: "",
			pos: 0,
			out: "",
		},
		{
			str: "hello world  ",
			pos: 10,
			out: "hello%20world",
		},
		{
			str: "helloworld",
			pos: 9,
			out: "helloworld",
		},
		{
			str: "hello world !!    ",
			pos: 13,
			out: "hello%20world%20!!",
		},
	}

	for _, test := range tests {
		t.Run(test.str+"_"+strconv.Itoa(test.pos), func(t *testing.T) {
			out := URLify(test.str, test.pos)
			if out != test.out {
				t.Errorf("actual=%q expected=%q", out, test.out)
			}
		})
	}
}
