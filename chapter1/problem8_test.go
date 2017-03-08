package chapter1

import (
	"reflect"
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	tests := map[string]struct {
		in   [][]int
		post [][]int
	}{
		"should set 0s if input has 0": {
			in: [][]int{
				[]int{1, 8, 6, 5, 0},
				[]int{0, 5, 8, 2, 2},
				[]int{3, 5, 3, 2, 2},
				[]int{6, 6, 4, 7, 3},
				[]int{8, 4, 3, 3, 1},
			},
			post: [][]int{
				[]int{0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0},
				[]int{0, 5, 3, 2, 0},
				[]int{0, 6, 4, 7, 0},
				[]int{0, 4, 3, 3, 0},
			},
		},

		"should not set if input doesn't have 0": {
			in: [][]int{
				[]int{1, 8, 6, 5, 1},
				[]int{1, 5, 8, 2, 2},
				[]int{3, 5, 3, 2, 2},
				[]int{6, 6, 4, 7, 3},
				[]int{8, 4, 3, 3, 1},
			},
			post: [][]int{
				[]int{1, 8, 6, 5, 1},
				[]int{1, 5, 8, 2, 2},
				[]int{3, 5, 3, 2, 2},
				[]int{6, 6, 4, 7, 3},
				[]int{8, 4, 3, 3, 1},
			},
		},

		"should not set if input is empty": {
			in:   [][]int{},
			post: [][]int{},
		},
	}

	for k, test := range tests {
		t.Run(k, func(t *testing.T) {
			ZeroMatrix(test.in)
			if !reflect.DeepEqual(test.in, test.post) {
				t.Errorf("actual=%v expected=%v", test.in, test.post)
			}
		})
	}
}
