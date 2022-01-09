package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const MAX_INT = int(^uint(0)>>1)
const MIN_INT = ^MAX_INT

func max(x, y int) int {
	if (x > y) {
		return x
	}
	return y
}

func oneSideMax(root *TreeNode, maxSum *int) int {
	if (root == nil ) {
		return 0;
	}

	var right = 0 
	var left = 0

	right = oneSideMax(root.Right, maxSum)
	left =  oneSideMax(root.Left, maxSum)

	sideMax := max(max(right + root.Val, left + root.Val), root.Val)
	
	*maxSum = max(*maxSum, max(sideMax, right + left + root.Val))

	return sideMax;
}

func maxPathSum(root *TreeNode) int {

	if (root == nil){
		return 0
	}
	var ans = MIN_INT

	oneSideMax(root, &ans)

	return ans
}
