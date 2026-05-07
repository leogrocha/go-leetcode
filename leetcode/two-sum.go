package leetcode

func TwoSum(nums []int, target int) []int {
	memory := make(map[int]int)
	for index, value := range nums {
		if v, ok := memory[target-value]; ok {
			return []int{v, index}
		}

		memory[value] = index
	}

	return []int{}
}
