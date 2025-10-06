package nce

import "slices"

func NotContainingElem(s1 []string, s2 []string) []string {
	var res []string

	for _, l := range s1 {
		if !slices.Contains(s2, l) {
			res = append(res, l)
		}
	}

	return res
}
