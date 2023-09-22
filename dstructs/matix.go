package dstructs

import "fmt"

func incidentMatrix() {
	num := 3
	sl1 := [][]int{}
	for i := 0; i < num; i++ {
		sl := make([]int, num)
		sl[i] = 1
		// fmt.Println(sl)
		sl1 = append(sl1, sl)
	}
	fmt.Println(sl1)
}
