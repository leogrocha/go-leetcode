package leetcode

import "sort"

func ContainsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	sort.Ints(nums)

	temp := nums[0]
	for index, item := range nums {
		if index != 0 {
			if item == temp {
				return true
			}
		}

		temp = item
	}

	return false
}
