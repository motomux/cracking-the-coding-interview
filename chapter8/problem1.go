package chapter8

// TripleStep calles TripleStepR to calculates possible ways to claim
// Time complexity O(3^n)
// Space complexity O(n)
func TripleStep(num int) int {
	if num == 0 {
		return 0
	}

	return TripleStepR(num)
}

// TripleStepR calculates possible ways to claim give number of stairs with 3, 2, 1 steps
func TripleStepR(num int) int {
	if num == 0 {
		return 1
	} else if num < 0 {
		return 0
	}

	return TripleStepR(num-1) + TripleStepR(num-2) + TripleStepR(num-3)
}

// TripleStep2 calles TripleStepR to calculates possible ways to claim
// Time complexity O(n)
// Space complexity O(n)
func TripleStep2(num int) int {
	if num == 0 {
		return 0
	}

	memo := make([]int, num+1)
	return TripleStepR2(num, memo)
}

// TripleStepR2 calculates possible ways to claim give number of stairs with 3, 2, 1 steps
func TripleStepR2(num int, memo []int) int {
	if num == 0 {
		return 1
	} else if num < 0 {
		return 0
	}

	if memo[num] == 0 {
		memo[num] = TripleStepR2(num-1, memo) +
			TripleStepR2(num-2, memo) +
			TripleStepR2(num-3, memo)
	}

	return memo[num]
}
