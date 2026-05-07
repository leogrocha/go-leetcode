package leetcode

import "slices"

func GetConcatenation(nums []int) []int {
	return slices.Concat(nums, slices.Clone(nums))
}
