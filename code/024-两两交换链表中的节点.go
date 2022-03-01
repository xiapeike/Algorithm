package main

// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
/*
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
*/
import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5}

	var l1head = new(ListNode)

	l1 := new(ListNode)
	l1.Val = -1
	l1head = l1

	for i := 0; i < len(nums1); i++ {
		var node = ListNode{Val: nums1[i]}
		l1.Next = &node
		l1 = l1.Next
	}

	result := swapPairs(l1head.Next)
	for result != nil {
		fmt.Printf("sum = %+v", *result)
		result = result.Next
	}

	//fmt.Printf("sum = %+v", *l1)
}
