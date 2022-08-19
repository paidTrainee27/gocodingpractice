package dstructs

import "fmt"

func oddArraysWithSingleLoop() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(arr1); i++ {
		if arr1[i]%2 == 0 {
			arr1 = append(arr1[:i], arr1[i+1:]...) //pop....removing even numbers
		}
	}
	fmt.Println(arr1)
}



func CrossMergeSliceWithHelper() {

	nums1 := []int{2, 4, 6, 8, 10}
	// m := 0
	nums2 := []int{1, 3, 5, 7, 9}
	leng := (len(nums2) + len(nums1))
	result := make([]int, leng)

	for i, j, k := 0, 0, 0; i < leng; i++ {
		if i%2 == 0 {
			result[i] = nums2[j]
			j++
		} else {
			result[i] = nums1[k]
			k++
		}
	}
	fmt.Println(result)

}
func CrossMergeSlice() {

	nums1 := []int{2, 4, 6, 8, 10}
	
	nums2 := []int{1, 3, 5, 7, 9}
	//12345678910
	leng := (len(nums2) + len(nums1))

	for i, j := 0, 0; i < leng; i++ {
		//only odd indices
		//123579
		if i%2 != 0 && i != 0 {
			nums2 = append(nums2[:i], append([]int{nums1[j]}, nums2[i:]...)...)
			j++
		}
	}
	fmt.Println(nums2)
}
