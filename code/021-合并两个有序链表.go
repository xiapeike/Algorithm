package main

// https://leetcode-cn.com/problems/merge-two-sorted-lists/
/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/
import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针做法 空间 o(m + n)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var index = new(ListNode)

	returnList := index
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			var node = new(ListNode)
			node.Val = list1.Val
			node.Next = nil
			index.Next = node
			index = index.Next
			list1 = list1.Next
		} else {
			var node = new(ListNode)
			node.Val = list2.Val
			node.Next = nil
			index.Next = node
			index = index.Next
			list2 = list2.Next
		}
	}

	if list1 != nil {
		index.Next = list1
	}

	if list2 != nil {
		index.Next = list2
	}

	return returnList.Next
}

/*递归做法 空间 o(1)*/
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	// 出口
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	// 递归函数
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists2(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists2(list1, list2.Next)
		return list2
	}
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
	result := mergeTwoLists2(l1head.Next, l2head.Next)
	for result != nil {
		fmt.Printf("sum = %+v", *result)
		result = result.Next
	}

	//fmt.Printf("sum = %+v", *l1)
}
