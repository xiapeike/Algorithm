package main

import (
	"code/util"
	"fmt"
)

// https://leetcode.cn/problems/binary-tree-inorder-traversal/
/*
给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
*/

/*
递归
*/

func inorderTraversal094(root *util.TreeNode) []int {
	var ans []int
	var mid_trace func(*util.TreeNode)

	mid_trace = func(node *util.TreeNode) {
		if node != nil {
			if node.Left != nil {
				mid_trace(node.Left)
			}

			ans = append(ans, node.Val)

			if node.Right != nil {
				mid_trace(node.Right)
			}
		}
	}

	mid_trace(root)
	return ans
}

/*迭代*/
func inorderTraversal(root *util.TreeNode) (res []int) {
	stack := []*util.TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}

func main() {
	input := []int{1, -1, 2, 3}
	//input := []int{2, 4}
	//input1 := "AB"
	result := inorderTraversal094(util.InitTree(input))
	fmt.Printf("result = %v", result)
}
