package main

import("fmt")

func main() {
	fmt.Println(maxScore("011101"))
}

func maxScore(s string) int {
	ones := 0
	for _, val := range s {
		if (val == '1') {
			ones ++
		}
	}
	leftones := 0
	result := 0
	for index, val := range s {
		if val == '1' {
			leftones ++
		}
		if index != len(s) - 1 {
			leftscore := index - leftones + 1
			rightscore := ones - leftones
			result = max(leftscore + rightscore, result)
		}
	}
	return result;
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}