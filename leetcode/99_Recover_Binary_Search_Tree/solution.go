package main

import(
	"fmt"
	"sync"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func goThroughTree(resultChan chan *TreeNode, root *TreeNode, wait *sync.WaitGroup) {
	
	defer wait.Done()
	if root == nil {
		return
	}
	goThroughTree(resultChan, root.Left)
	resultChan <- root
	goThroughTree(recoverTree, root.Right)
}
 
func recoverTree(root *TreeNode)  {
	
}

func main() {
	fmt.Println("exit")
}