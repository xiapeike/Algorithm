package main

// https://leetcode-cn.com/problems/add-two-numbers/
import (
	"code/util"
	"fmt"
)

// ListNode define

func addTwoNumbers1(l1 *util.ListNode, l2 *util.ListNode) *util.ListNode {
	var returnList = new(util.ListNode)
	head := returnList
	// 分别得到两个链表长度
	var l1len, l2len, shortlen int = 0, 0, 0
	l1head, l2head := l1, l2
	for l1head != nil {
		l1len++
		l1head = l1head.Next
	}
	for l2head != nil {
		l2len++
		l2head = l2head.Next
	}
	shortlen = l1len
	if l1len > l2len {
		shortlen = l2len
	}

	jinwei := 0
	for i := 1; i <= shortlen; i++ {
		add1 := l1.Val
		add2 := l2.Val
		newVar := add1 + add2
		if jinwei > 0 {
			newVar++
			jinwei = 0
		}
		if newVar >= 10 {
			jinwei = 1
			newVar = newVar % 10
		}
		node := new(util.ListNode)
		node.Val = newVar
		returnList.Next = node
		returnList = returnList.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		newVar := l1.Val
		if jinwei > 0 {
			newVar++
			jinwei = 0
		}
		if newVar >= 10 {
			jinwei = 1
			newVar = newVar % 10
		}
		node := new(util.ListNode)
		node.Val = newVar
		returnList.Next = node
		returnList = returnList.Next
		l1 = l1.Next
	}

	for l2 != nil {
		newVar := l2.Val
		if jinwei > 0 {
			newVar++
			jinwei = 0
		}
		if newVar >= 10 {
			jinwei = 1
			newVar = newVar % 10
		}
		node := new(util.ListNode)
		node.Val = newVar
		returnList.Next = node
		returnList = returnList.Next
		l2 = l2.Next
	}
	if jinwei > 0 {
		node := new(util.ListNode)
		node.Val = jinwei
		returnList.Next = node
		returnList = returnList.Next
	}

	return head.Next
}

func main() {
	nums1 := []int{0}
	nums2 := []int{0}
	var l1head = new(util.ListNode)
	var l2head = new(util.ListNode)
	l1 := new(util.ListNode)
	l1.Val = -1
	l1head = l1
	l2 := new(util.ListNode)
	l2.Val = -1
	l2head = l2
	for i := 0; i < len(nums1); i++ {
		var node = util.ListNode{Val: nums1[i]}
		l1.Next = &node
		l1 = l1.Next
	}

	for i := 0; i < len(nums2); i++ {
		var node = util.ListNode{Val: nums2[i]}
		l2.Next = &node
		l2 = l2.Next
	}

	result := addTwoNumbers1(l1head.Next, l2head.Next)
	for result != nil {
		fmt.Printf("sum = %+v", *result)
		result = result.Next
	}

	//fmt.Printf("sum = %+v", *l1)
}
