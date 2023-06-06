package main

import (
	"code/util"
	"fmt"
)

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
/*
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。
*/

/*

 */

func maxDepth104(root *util.TreeNode) int {

	if root == nil {
		return 0
	}
	level := 0 // 层数
	var queue []*util.TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		item := len(queue)
		for i := 0; i < item; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		level = level + 1
	}

	return level
}
func main() {
	//input := []int{1, 2, 2, 3, 4, 4, 3}
	//input := []int{1, 2, 2, -1, 3, -1, 3}
	//input := []int{1, 2, 2, 2, -1, 2}
	input := []int{1, 2, 3}

	result := maxDepth104(util.InitTree(input))
	fmt.Printf("result = %v", result)
}
