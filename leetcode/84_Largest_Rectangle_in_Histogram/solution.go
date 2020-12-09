package leetcode

type rectangle struct {
	location int
	height   int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxRectangle(rectangles []rectangle, len int) int {
	result := 0
	for i := len - 1; i > 0; i-- {
		result = max(result, rectangles[i].height*(rectangles[len-1].location-rectangles[i-1].location))
	}
	return result
}

func largestRectangleArea(heights []int) int {
	rectangles := make([]rectangle, 0)
	rectangles = append(rectangles, rectangle{-1, 0})
	recLen := 1
	result := 0
	for index := range heights {
		for tail := recLen - 1; tail >= 0 && rectangles[tail].height >= heights[index]; tail-- {
			result = max(result, getMaxRectangle(rectangles, recLen))
			recLen--
		}
		if recLen == len(rectangles) {
			rectangles = append(rectangles, rectangle{index, heights[index]})
		} else {
			rectangles[recLen] = rectangle{index, heights[index]}
		}
		recLen++
	}
	return max(result, getMaxRectangle(rectangles, recLen))
}
