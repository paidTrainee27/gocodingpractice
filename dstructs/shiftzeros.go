package dstructs

import "fmt"

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//
// Example:
// Input: nums = [0,3,0,1,12]
// Output: [3,1,12,0,0]

func ShiftZerosToEnd() {
	nums := []int{0, 3, 0, 1, 12}
	siz := len(nums)
	i := 0
	for i < siz {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
			siz--
		} else {
			i++
		}
	}

	fmt.Println(nums)
}
