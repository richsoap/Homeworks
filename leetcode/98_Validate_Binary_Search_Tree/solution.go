package leetcode

import "fmt"

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	resChan := make(chan int)
	result := make([]int, 0)
	go goThrough(root, true, resChan)
	for val := range resChan {
		result = append(result, val)
	}
	prev := result[0] - 1
	fmt.Printf("%v", result)
	for _, val := range result {
		if val <= prev {
			return false
		}
		prev = val
	}
	return true
}

func goThrough(root *TreeNode, isRoot bool, resChan chan int) {
	if isRoot {
		defer close(resChan)
	}
	if root == nil {
		return
	}
	goThrough(root.Left, false, resChan)
	resChan <- root.Val
	goThrough(root.Right, false, resChan)
}
