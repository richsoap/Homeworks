package leetcode

func trap(height []int) int {
	maxHeight := 0
	leftSpace := 0
	wallSpace := 0
	rightSpace := 0
	for index := range height {
		wallSpace += height[index]
		if height[index] > maxHeight {
			leftSpace += (height[index] - maxHeight) * index
			maxHeight = height[index]
		}
	}
	allSpace := maxHeight * len(height)
	maxHeight = 0
	for index := len(height) - 1; index >= 0; index-- {
		if height[index] > maxHeight {
			rightSpace += (height[index] - maxHeight) * (len(height) - 1 - index)
			maxHeight = height[index]
		}
	}
	return allSpace - leftSpace - rightSpace - wallSpace
}
