package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// Check if subtrees have height within 1. If they do, check if the
	// subtrees are balanced
	return ((Abs(height(root.Left) - height(root.Right))) < 2) && isBalanced(root.Right) && isBalanced(root.Left)
}

func height(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return 1 + max(height(root.Left), height(root.Right))
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
