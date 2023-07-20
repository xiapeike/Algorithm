package main

import (
	"code/util"
)

// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/
/*
给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。
*/

/*
方法一：前序遍历
将二叉树展开为单链表之后，单链表中的节点顺序即为二叉树的前序遍历访问各节点的顺序。
因此，可以对二叉树进行前序遍历，获得各节点被访问到的顺序。由于将二叉树展开为链表之后会破坏二叉树的结构，
因此在前序遍历结束之后更新每个节点的左右子节点的信息，将二叉树展开为单链表。

*/

func flatten114(root *util.TreeNode) {
	list := preorderTraversal(root)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}
func preorderTraversal(root *util.TreeNode) []*util.TreeNode {
	list := []*util.TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, preorderTraversal(root.Left)...)
		list = append(list, preorderTraversal(root.Right)...)
	}
	return list
}

/*
方法三：寻找前驱节点

注意到前序遍历访问各节点的顺序是根节点、左子树、右子树。如果一个节点的左子节点为空，则该节点不需要进行展开操作。

如果一个节点的左子节点不为空，则该节点的左子树中的最后一个节点被访问之后，该节点的右子节点被访问。

该节点的左子树中最后一个被访问的节点是左子树中的最右边的节点，也是该节点的前驱节点。因此，问题转化成寻找当前节点的前驱节点。

具体做法是，对于当前节点，如果其左子节点不为空，则在其左子树中找到最右边的节点，作为前驱节点，将当前节点的右子节点赋给前驱节点的右子节点，

然后将当前节点的左子节点赋给当前节点的右子节点，并将当前节点的左子节点设为空。对当前节点处理结束后，继续处理链表中的下一个节点，直到所有节点都处理结束。
*/
func flatten1142(root *util.TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			predecessor := next
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			predecessor.Right = curr.Right
			curr.Left, curr.Right = nil, next
		}
		curr = curr.Right
	}
}

func main() {
	input := []int{1, 2, 5, 3, 4, -1, 6}

	flatten114(util.InitTree(input))
}
