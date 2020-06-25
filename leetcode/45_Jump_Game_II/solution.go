package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 空输入的定义？
// 到0是不是0？

func jump(nums []int) int {
	minSteps := make([]int, len(nums), len(nums))
	for index := range minSteps {
		minSteps[index] = len(nums)
	}
	minSteps[0] = 0
	for index := range minSteps {
		for i := min(nums[index], len(nums)-index-1); i > 0 && minSteps[index+i] > minSteps[index]+1; i-- {
			minSteps[index+i] = minSteps[index] + 1
		}
	}
	return minSteps[len(minSteps)-1]
}
