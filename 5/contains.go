package contains

import "slices"

func Contains(s1 []int, s2 []int) ([]int, bool) {
	var res []int
	status := false

	for _, l := range s1 {
		if slices.Contains(s2, l) {
			res = append(res, l)
			status = true
		}
	}

	return res, status
}
