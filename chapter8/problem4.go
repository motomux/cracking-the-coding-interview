package chapter8

// PowerSet returns all subsets of given set
func PowerSet(set string) []string {
	if len(set) == 0 {
		return []string{""}
	}

	var res []string
	for _, subset := range PowerSet(set[1:]) {
		res = append(res, subset, string(set[0])+subset)
	}

	return res
}
