package main

import (
	"fmt"
	"slices"
)

func main() {
	// var nums = []int{1, 2, 3, 4, 5}

	// nums := make([]int, 5)
	// nums[0] = 1
	// fmt.Println(nums)
	// fmt.Println(nums == nil)
	// fmt.Println(cap(nums))
	// nums = append(nums, 6)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	//copy function

	// var nums2 = make([]int, 0, 5)
	// nums2 = append(nums2, 2)
	// var nums3 = make([]int, len(nums2))

	// //copy
	// copy(nums3, nums2)

	// fmt.Println(nums2, nums3)

	// //slice operator

	// nums := []int{0, 1, 2, 3, 4, 5}
	// fmt.Println(nums[:3])

	//slice package

	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	fmt.Println(slices.Equal(nums1, nums2))

}
