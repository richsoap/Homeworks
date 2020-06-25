package leetcode

func findDiagonalOrder(nums [][]int) []int {
	edge := 0
	bottom := len(nums)
	count := 0
	for index := range nums {
		count += len(nums[index])
		edge = max(edge, len(nums[index]))
	}
	result := make([]int, 0, count)
	jumprec := make([]int, len(nums), len(nums))
	for index := range jumprec {
		jumprec[index] = index - 1
	}
	for y := 0;y < bottom;y ++ {
		result = append(result, travel(0, y, edge, &nums, &jumprec)...)
	}
	for x := 1;x < edge; x ++ {
		result = append(result, travel(x, len(nums)-1, edge, &nums, &jumprec)...)
	}
	return result
}


func travel(x, y, edge int, nums *[][]int, jumprec *[]int) []int {
	result := make([]int, 0)
	for {
		result = append(result, (*nums)[x][y])
		nline := (*jumprec)[y]
		for nline >= 0 {
			if len((*nums)[nline]) <= y-nline+x {
				nline = (*jumprec)[nline]
			} else {
				break
			}
		}
		if nline < 0 {
			break
		}
		(*jumprec)[y] = nline
		x = x + y - nline
		y = nline
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
