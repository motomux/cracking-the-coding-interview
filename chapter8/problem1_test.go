package chapter8

import (
	"strconv"
	"testing"
)

func TestTripleStep(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{
			in:  0,
			out: 0,
		},
		{
			in:  1,
			out: 1,
		},
		{
			in:  2,
			out: 2,
		},
		{
			in:  3,
			out: 4,
		},
		{
			in:  5,
			out: 13,
		},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.in), func(t *testing.T) {
			out := TripleStep(test.in)
			if out != test.out {
				t.Errorf("actual=%d expected=%d", out, test.out)
			}
		})
	}
}

func TestTripleStep2(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{
			in:  0,
			out: 0,
		},
		{
			in:  1,
			out: 1,
		},
		{
			in:  2,
			out: 2,
		},
		{
			in:  3,
			out: 4,
		},
		{
			in:  5,
			out: 13,
		},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.in), func(t *testing.T) {
			out := TripleStep2(test.in)
			if out != test.out {
				t.Errorf("actual=%d expected=%d", out, test.out)
			}
		})
	}
}
