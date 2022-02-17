package main

// https://leetcode-cn.com/problems/merge-k-sorted-lists/
/*
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。
*/
import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 查找所有链表第一个值最小的索引
func getMinIndex(lists []*ListNode) int {
	index := -1
	minValue := math.MaxInt8
	for i, node := range lists {
		if node != nil && node.Val < minValue {
			index = i
			minValue = node.Val
		}
	}
	return index
}

// k指针做法 空间 o(m * n)
func mergeKLists(lists []*ListNode) *ListNode {
	var index = getMinIndex(lists)

	if index == -1 {
		return nil
	}
	head := lists[index]
	lists[index] = lists[index].Next
	head.Next = mergeKLists(lists)
	return head
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{1, 3, 4}

	var l1head = new(ListNode)
	var l2head = new(ListNode)
	l1 := new(ListNode)
	l1.Val = -1
	l1head = l1

	for i := 0; i < len(nums1); i++ {
		var node = ListNode{Val: nums1[i]}
		l1.Next = &node
		l1 = l1.Next
	}
	l2 := new(ListNode)
	l2.Val = -1
	l2head = l2

	for i := 0; i < len(nums2); i++ {
		var node = ListNode{Val: nums2[i]}
		l2.Next = &node
		l2 = l2.Next
	}

	lists := []*ListNode{l1head.Next, l2head.Next}
	result := mergeKLists(lists)
	for result != nil {
		fmt.Printf("sum = %+v", *result)
		result = result.Next
	}

	//fmt.Printf("sum = %+v", *l1)
}
