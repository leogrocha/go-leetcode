package leetcode

import "sort"

func SingleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sort.Ints(nums)

	var containsDuplicate bool

	var index int
	for {
		if index != len(nums)-1 {
			if nums[index] == nums[index+1] {
				containsDuplicate = true
			} else if nums[index] != nums[index+1] && containsDuplicate {
				containsDuplicate = false
			} else if nums[index] != nums[index+1] && !containsDuplicate {
				return nums[index]
			}
		} else {
			return nums[index]
		}
		index++
	}

}
