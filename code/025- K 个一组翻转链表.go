package main

// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
/*
给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。

k是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。

进阶：

你可以设计一个只使用常数额外空间的算法来解决此问题吗？
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

*/

import (
	"code/util"
	"fmt"
)

// 采用递归的做法
func reverseKGroup(head *util.ListNode, k int) *util.ListNode {
	// 找到第k+1节点
	knode := head
	if k == 1 {
		return head
	}
	for i := 0; i < k-1; i++ {
		if knode == nil || knode.Next == nil {
			return head
		}
		knode = knode.Next
	}

	newHead := head.Next

	head.Next = reverseKGroup(knode.Next, k)

	// 翻转k-1的结点
	var swapKnodes func(head *util.ListNode, k int) *util.ListNode
	swapKnodes = func(head *util.ListNode, k int) *util.ListNode {
		resultNode := &util.ListNode{}
		// 采用头插法逆转链表
		for i := 0; i < k; i++ {
			newHead := head.Next
			head.Next = resultNode.Next
			resultNode.Next = head
			head = newHead
		}
		return resultNode.Next
	}
	knode = swapKnodes(newHead, k-1)

	newHead.Next = head
	return knode
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5}

	var l1head = new(util.ListNode)

	l1 := new(util.ListNode)
	l1.Val = -1
	l1head = l1

	for i := 0; i < len(nums1); i++ {
		var node = util.ListNode{Val: nums1[i]}
		l1.Next = &node
		l1 = l1.Next
	}

	result := reverseKGroup(l1head.Next, 1)
	for result != nil {
		fmt.Printf("sum = %+v", *result)
		result = result.Next
	}

	//fmt.Printf("sum = %+v", *l1)
}
