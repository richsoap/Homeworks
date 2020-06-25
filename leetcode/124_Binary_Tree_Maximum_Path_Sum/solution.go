package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(nums ...int) int {
	if len(nums) <= 0 {
		return 0
	}
	result := nums[0]
	for _, val := range nums {
		if result < val {
			result = val
		}
	}
	return result
}

func maxPathSum(root *TreeNode) int {
	result := root.Val
	goThroughTree(root, &result)
	return result
}

// return maxPath and maxResult
func goThroughTree(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	leftPath := goThroughTree(root.Left, result)
	rightPath := goThroughTree(root.Right, result)
	pathResult := max(0, leftPath, rightPath) + root.Val
	*result = max(*result, leftPath+root.Val, rightPath+root.Val, rightPath+leftPath+root.Val, root.Val)
	return pathResult
}
