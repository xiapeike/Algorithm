package util

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InitTree(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0]}
	queue := list.New()
	queue.PushBack(root)

	i := 1
	for i < n {
		parent := queue.Remove(queue.Front()).(*TreeNode)

		leftVal, rightVal := -1, -1
		if i < n {
			leftVal = nums[i]
		}
		if i+1 < n {
			rightVal = nums[i+1]
		}

		if leftVal != -1 {
			leftNode := &TreeNode{Val: leftVal}
			parent.Left = leftNode
			queue.PushBack(leftNode)
		}
		if rightVal != -1 {
			rightNode := &TreeNode{Val: rightVal}
			parent.Right = rightNode
			queue.PushBack(rightNode)
		}

		i += 2
	}

	return root
}
