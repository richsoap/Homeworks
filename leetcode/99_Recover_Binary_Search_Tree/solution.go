package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goThroughTree(resultChan chan *TreeNode, root *TreeNode) {
	if root == nil {
		return
	}
	goThroughTree(resultChan, root.Left)
	resultChan <- root
	goThroughTree(resultChan, root.Right)
}

func recoverTree(root *TreeNode) {
	resultChan := make(chan *TreeNode)
	var prevNode, highNode, lowNode *TreeNode
	prevNode = nil
	highNode = nil
	lowNode = nil
	go func() {
		goThroughTree(resultChan, root)
		close(resultChan)
	}()

	for node := range resultChan {
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
