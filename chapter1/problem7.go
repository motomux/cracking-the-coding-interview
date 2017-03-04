package chapter1

// RotateMatrix rotates n x n matrix of int
func RotateMatrix(m [][]int) {
	for i := 0; i < (len(m)+1)/2; i++ {
		s, e := i, len(m)-i-1
		for k := 0; k < e-s; k++ {
			m[s][s+k], m[s+k][e], m[e][e-k], m[e-k][s] =
				m[e-k][s], m[s][s+k], m[s+k][e], m[e][e-k]
		}
	}
}
