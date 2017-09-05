package chapter8

import (
	"reflect"
	"testing"
)

func TestPowerSet(t *testing.T) {
	tests := []struct {
		in  string
		out []string
	}{
		{
			in:  "",
			out: []string{""},
		},
		{
			in:  "a",
			out: []string{"", "a"},
		},
		{
			in:  "ab",
			out: []string{"", "a", "b", "ab"},
		},
		{
			in:  "abc",
			out: []string{"", "a", "b", "ab", "c", "ac", "bc", "abc"},
		},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			out := PowerSet(test.in)
			if !reflect.DeepEqual(out, test.out) {
				t.Errorf("actual=%q expected=%q", out, test.out)
			}
		})
	}
}
