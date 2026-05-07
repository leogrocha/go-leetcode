package leetcode

import "slices"

func MoveZeroes(nums []int) {
	startLenght := len(nums)
	nums = slices.DeleteFunc(nums, func(v int) bool { return v == 0 })

	for startLenght != len(nums) {
		nums = append(nums, 0)
	}
}
