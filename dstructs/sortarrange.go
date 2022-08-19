package dstructs


func findLargest(n []int) int {
	larg := 0 // won't require this if n is array
	for j := 1; j < len(n); j++ {
		if larg < n[j] {
			larg = n[j]
		}

	}
	return larg
}
