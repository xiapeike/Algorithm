package main

import (
	"code/util"
	"fmt"
)

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
/*
给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
*/

/*
1. 取先序遍历第一个结点，即为根结点
2. 根据第一个结点，将中序数组分为左右两部分
3. 根据左右两部分的数量，将先序数组也分为对应的左右两部分
4. 递归地根据左部分构建左子树
5. 递归地根据右部分构建右子树
*/

func buildTree105(preorder []int, inorder []int) *util.TreeNode {
	if len(preorder) == 1 {
		if preorder[0] == -1 {
			return nil
		} else {
			return &util.TreeNode{Val: preorder[0]}
		}
	}

	root := &util.TreeNode{Val: preorder[0]}
	inorder_left_part, inorder_right_part := create_inorder_parts(preorder[0], inorder)
	if len(inorder_left_part) > 0 {
		root.Left = buildTree105(preorder[1:len(inorder_left_part)+1], inorder_left_part)
	}

	if len(inorder_right_part) > 0 {
		root.Right = buildTree105(preorder[len(inorder_left_part)+1:], inorder_right_part)
	}
	return root
}

func create_inorder_parts(root_value int, inorder []int) ([]int, []int) {
	var left_part, right_part []int
	if len(inorder) == 0 {
		return left_part, right_part
	}
	index := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root_value {
			index = i
			break
		}
	}
	for i := 0; i < index; i++ {
		left_part = append(left_part, inorder[i])
	}

	for i := index + 1; i < len(inorder); i++ {
		right_part = append(right_part, inorder[i])
	}
	return left_part, right_part

}
func main() {
	input := []int{3, 9, 20, 15, 7}
	input2 := []int{9, 3, 15, 20, 7}

	result := buildTree105(input, input2)
	fmt.Printf("result = %v", result)
}
