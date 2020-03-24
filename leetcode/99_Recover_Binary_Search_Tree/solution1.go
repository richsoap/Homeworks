package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goThroughTree(array *[]*TreeNode, root *TreeNode) {
	if root == nil {
		return
	}
	goThroughTree(array, root.Left)
	*array = append(*array, root)
	goThroughTree(array, root.Right)
}

func recoverTree(root *TreeNode) {
	var prevNode, highNode, lowNode *TreeNode
	nodeArray := make([]*TreeNode, 0)
	prevNode = nil
	highNode = nil
	lowNode = nil

	goThroughTree(&nodeArray, root)

	for _, node := range nodeArray {
		if prevNode != nil && node.Val < prevNode.Val {
			if highNode == nil {
				highNode = prevNode
			}
			lowNode = node
		}
		prevNode = node
	}
	highNode.Val, lowNode.Val = lowNode.Val, highNode.Val
}

func main() {
	fmt.Println("exit")
}

//20ms????
