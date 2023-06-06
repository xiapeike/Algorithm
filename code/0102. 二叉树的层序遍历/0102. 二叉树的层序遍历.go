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

 */

func levelOrder102(root *util.TreeNode) [][]int {
	var ans [][]int
	var queue []*util.TreeNode
	var lever_array []int
	if root != nil {
		queue = append(queue, root)
		level := 1 // 每层结点计数
		lever_array = append(lever_array, level)

		level_item := []int{}
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			cur_level := lever_array[0]
			lever_array = lever_array[1:]

			if cur_level > level {
				ans = append(ans, append([]int{}, level_item...))
				level = cur_level
				level_item = []int{}
			}
			level_item = append(level_item, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
				lever_array = append(lever_array, cur_level+1)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
				lever_array = append(lever_array, cur_level+1)
			}
		}
		if len(level_item) > 0 {
			ans = append(ans, append([]int{}, level_item...))
		}
	}
	return ans
}

func main() {
	//input := []int{1, 2, 2, 3, 4, 4, 3}
	//input := []int{1, 2, 2, -1, 3, -1, 3}
	//input := []int{1, 2, 2, 2, -1, 2}
	input := []int{1, 2, 3}

	result := levelOrder102(util.InitTree(input))
	fmt.Printf("result = %v", result)
}
