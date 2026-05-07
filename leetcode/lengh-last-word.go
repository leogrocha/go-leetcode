package leetcode

import "strings"

func LengthOfLastWord(s string) int {
	words := strings.Split(s, " ")
	for index := len(words) - 1; index >= 0; index-- {
		if len(words[index]) != 0 {
			return len(words[index])
		}
	}
	return 0
}
