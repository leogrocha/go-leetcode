package leetcode

func MajorityElement(nums []int) int {
	mem := make(map[int]int)

	for _, item := range nums {
		if v, exits := mem[item]; exits {
			mem[item] = v + 1
		} else {
			mem[item] = 1
		}
	}

	var result int
	var bigger int
	for key, value := range mem {
		if value > bigger {
			bigger = value
			result = key
		}
	}

	return result
}
