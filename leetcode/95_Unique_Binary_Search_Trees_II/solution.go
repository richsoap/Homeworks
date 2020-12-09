package main

import(
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func makeTrees(start, end int) []*TreeNode {
	result := make([]*TreeNode, 0)
	for i := start;i < end;i ++ {
		leftTrees := makeTrees(start,i)
		rightTrees := makeTrees(i + 1, end)
		if len(leftTrees) == 0 {
			leftTrees = append(leftTrees, nil)
		}
		if len(rightTrees) == 0 {
			rightTrees = append(rightTrees, nil)
		}
		for _, leftNode := range(leftTrees) {
			for _, rightNode := range(rightTrees) {
				result = append(result, &TreeNode{i, leftNode, rightNode})
			}
		}
	}
	return result
}

func generateTrees(n int) []*TreeNode {
	return makeTrees(1, n + 1)
}

func main() {
	result := generateTrees(3)
	fmt.Println(len(result))
}