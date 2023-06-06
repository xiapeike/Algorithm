package main

import (
	"code/util"
	"fmt"
)

// https://leetcode.cn/problems/validate-binary-search-tree/
/*
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
*/

/*
先序遍历得到一个数组
判断该数组是否是对称的
*/

func isValidBST098(root *util.TreeNode) bool {
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

	result := true
	for i := 0; i < len(ans)-1; i++ {
		if ans[i] >= ans[i+1] {
			result = false
			break
		}
	}
	return result
}

func main() {
	//input := []int{1, 2, 2, 3, 4, 4, 3}
	//input := []int{1, 2, 2, -1, 3, -1, 3}
	//input := []int{1, 2, 2, 2, -1, 2}
	input := []int{1, 2, 3}

	result := isValidBST098(util.InitTree(input))
	fmt.Printf("result = %v", result)
}
