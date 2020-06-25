package leetcode

func firstMissingPositive(nums []int) int {
	showRecord := make(map[int]int)
	for _, val := range nums {
		if val > 0 {
			showRecord[val] = 1
		}
	}
	result := 1
	for {
		if _, existed := showRecord[result]; !existed {
			break
		}
		result++
	}
	return result
}
