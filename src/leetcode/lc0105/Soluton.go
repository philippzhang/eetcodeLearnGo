package lc0105

import (
	"leetcode/structures"
)

type TreeNode = structures.TreeNode

func helper(preStart int, preEnd int, inStart int, inEnd int, preorder []int, inorder []int) *TreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}
	root := &TreeNode{
		Val: preorder[preStart],
	}

	inRoot := inStart
	for preorder[preStart] != inorder[inRoot] {
		inRoot++
	}
	len := inRoot - inStart

	root.Left = helper(preStart+1, preStart+len, inStart, inRoot-1, preorder, inorder)
	root.Right = helper(preStart+len+1, preEnd, inRoot+1, inEnd, preorder, inorder)

	return root
}
func buildTree(preorder []int, inorder []int) *TreeNode {
	return helper(0, len(preorder)-1, 0, len(inorder)-1, preorder, inorder)
}
