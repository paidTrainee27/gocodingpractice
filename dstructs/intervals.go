package dstructs

import "fmt"

func mergeIntervals() {

	iv := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	final := make([][]int, 0)
	m1 := iv[0]

	for i := 1; i < len(iv); i++ {
		z := iv[i]
		if z[0] < m1[0] || z[0] > m1[1] {
			m1[0] = z[0]
		}
		if z[1] > m1[1] {
			m1[1] = z[1]
		}

		// final = append(final, []int{m1[0], m1[1]})
		final = appendSl(final, m1)

	}

	fmt.Println(final)
}

func appendSl(sl [][]int, s []int) [][]int {
	s1 := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		s1[i] = s[i]
	}
	sl = append(sl, s1)

	return sl
}
