package main

import("fmt")

func main() {
	input := []int{1,2,3,4,5,6,1}
	fmt.Println(maxScore(input, 3))
}

func maxScore(cardPoints []int, k int) int {
	record := make([]int, len(cardPoints), len(cardPoints))
	for index := range cardPoints {
		record[index] = cardPoints[index]
		if index != 0 {
			record[index] += record[index - 1]
		}
	}
	if k == len(record) {
		return record[k - 1]
	}
	result := record[len(record) - 1] - record[len(record) - k - 1]
	for index := 0; index < k;index ++ {
		result = max(result, record[index] + record[len(record) - 1] - record[len(record) - k + index])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}