package main

// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
*/
import (
	"code/util"
	"fmt"
)

func removeNthFromEnd(head *util.ListNode, n int) *util.ListNode {
	var returnList = new(util.ListNode)
	result := returnList
	result.Next = head

	len_list := 0 // 用于保存当前列表的总长度
	for head != nil {
		len_list++
		head = head.Next
	}
	if len_list < n {
		return new(util.ListNode)
	}
	del_index := len_list - n
	head = result

	for ; del_index > 0; del_index-- {
		head = head.Next
	}

	head.Next = head.Next.Next

	return result.Next
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

	result := removeNthFromEnd(l1head.Next, 5)
	for result != nil {
		fmt.Printf("sum = %+v", *result)
		result = result.Next
	}

	//fmt.Printf("sum = %+v", *l1)
}
