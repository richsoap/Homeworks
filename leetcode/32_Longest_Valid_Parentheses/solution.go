package leetcode

func longestValidParentheses(s string) int {
	record := make([]int, len(s), len(s))
	for endIndex := range record {
		record[endIndex] = 0
		if s[endIndex] == '(' || endIndex == 0 {
			continue
		}
		pairIndex := endIndex - record[endIndex-1] - 1
		if pairIndex >= 0 && s[pairIndex] == '(' {
			record[endIndex] = record[endIndex-1] + 2
			if pairIndex > 0 {
				record[endIndex] += record[pairIndex-1]
			}
		}
	}
	result := 0
	for _, val := range record {
		if result < val {
			result = val
		}
	}
	return result
}
