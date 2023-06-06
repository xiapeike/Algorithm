package main

import (
	"code/util"
	"fmt"
)

// https://leetcode.cn/problems/symmetric-tree/
/*
给你一个二叉树的根节点 root ， 检查它是否轴对称。
*/

/*
先序遍历得到一个数组
判断该数组是否是对称的
*/

func isSymmetric0101(root *util.TreeNode) bool {

	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	var dps func(*util.TreeNode, *util.TreeNode) bool
	dps = func(left *util.TreeNode, right *util.TreeNode) bool {
		if left != nil && right != nil {
			if left.Val != right.Val {
				return false
			} else {
				return dps(left.Left, right.Right) && dps(left.Right, right.Left)
			}

		} else if left != nil {
			return false
		} else if right != nil {
			return false
		} else {
			return true
		}
	}
	return dps(root.Left, root.Right)
}

func main() {
	//input := []int{1, 2, 2, 3, 4, 4, 3}
	//input := []int{1, 2, 2, -1, 3, -1, 3}
	//input := []int{1, 2, 2, 2, -1, 2}
	input := []int{1, 2, 3}

	result := isSymmetric0101(util.InitTree(input))
	fmt.Printf("result = %v", result)
}
