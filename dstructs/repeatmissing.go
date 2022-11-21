package dstructs

import (
	"fmt"
	"sort"
)

func FindFirstNonRepeatingChar() {
	strarr := []string{"a", "a", "b", "a", "c", "c", "c", "d", "e", "e", "f"}
	rest1 := make(map[string]int)

	for _, v := range strarr {
		key := string(v)
		if _, ok := rest1[key]; ok {
			fmt.Println("Found", key)
			break
		}
		rest1[key] += 1
	}

	fmt.Println(rest1)
}

func countRepeatedNumsWithoutMap() {
	str := []int{1, 3, 5, 10, 4, 5}
	sort.Ints(str)
	maxNum := str[len(str)-1] + 1

	lookUp := make([]int, maxNum)
	for i := 0; i < len(str); i++ {
		lookUp[str[i]]++
	}
	fmt.Println(lookUp)
}

func countRepeatedNumsMap() {
	str := []int{1, 3, 5, 10, 4, 5}
	lookUp := make(map[int]int)
	for i := 0; i < len(str); i++ {
		lookUp[str[i]]++
	}
	fmt.Println(lookUp)
}

func findMissing1() {
	intArr := []int{1, 2, 6, 7, 8, 9, 10, 11, 4, 5, 11}
	result := make(map[int]int)
	// max := findLargest(intArr)
	for k := 0; k < len(intArr); k++ {

		result[intArr[k]]++
		if _, ok := result[k]; !ok {
			result[k] = 0
		}
	}
	fmt.Println(result)

}
