package lc0124

import "math"

import (
	"leetcode/structures"
)

type TreeNode = structures.TreeNode

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := math.MinInt32
	getPathSum(root, &max)
	return max
}

func getPathSum(root *TreeNode, maxSum *int) int {
	if root == nil {
		return math.MinInt32
	}
	left := max(0, getPathSum(root.Left, maxSum))
	right := max(0, getPathSum(root.Right, maxSum))
	*maxSum = max(*maxSum, left+right+root.Val)
	return root.Val + max(left, right)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
