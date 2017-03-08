package chapter1

// ZeroMatrix set 0 if an element on the same row or column has 0
// Time complexity O(n^2)
// Space complexity O(n)
func ZeroMatrix(m [][]int) {
	if len(m) == 0 {
		return
	}

	rows := make([]bool, len(m))
	cols := make([]bool, len(m[0]))

	for row := 0; row < len(m); row++ {
		for col := 0; col < len(m[row]); col++ {
			if m[row][col] == 0 {
				rows[row] = true
				cols[col] = true
			}
		}
	}

	for row, flag := range rows {
		if flag {
			for i := 0; i < len(m[row]); i++ {
				m[row][i] = 0
			}
		}
	}

	for col, flag := range cols {
		if flag {
			for i := 0; i < len(m); i++ {
				m[i][col] = 0
			}
		}
	}
}
