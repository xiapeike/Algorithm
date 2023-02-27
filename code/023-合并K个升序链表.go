package main

// https://leetcode-cn.com/problems/merge-k-sorted-lists/
/*
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。
*/
import (
	"code/util"
	"container/heap"
	"fmt"
	"math"
)

// 查找所有链表第一个值最小的索引
func getMinIndex(lists []*util.ListNode) int {
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

// 递归做法
func mergeKLists(lists []*util.ListNode) *ListNode {
	var index = getMinIndex(lists)

	if index == -1 {
		return nil
	}
	head := lists[index]
	lists[index] = lists[index].Next
	head.Next = mergeKLists(lists)
	return head
}

// 两两合并做法
func mergeKLists2(lists []*ListNode) *ListNode {
	var pre, cur *ListNode
	n := len(lists)
	for i := 0; i < n; i++ {
		if i == 0 {
			pre = lists[i]
			continue
		}
		cur = lists[i]
		pre = merge(pre, cur)
	}
	return pre
}

func merge(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				cur.Next = l1
				l1 = l1.Next
			} else {
				cur.Next = l2
				l2 = l2.Next
			}
			cur = cur.Next
		} else if l1 != nil {
			cur.Next = l1
			break
		} else {
			cur.Next = l2
			break
		}
	}
	return head.Next
}

type minHeap []*ListNode

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

// 所有类型都实现了空接口。这意味着，如果您编写一个函数以 interface{} 值作为参数，那么您可以为该函数提供任何值。
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 小顶堆做法
func mergeKLists3(lists []*ListNode) *ListNode {
	h := new(minHeap)
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}

	dummyHead := new(ListNode)
	pre := dummyHead
	for h.Len() > 0 {
		tmp := heap.Pop(h).(*ListNode)
		if tmp.Next != nil {
			heap.Push(h, tmp.Next)
		}
		pre.Next = tmp
		pre = pre.Next
	}

	return dummyHead.Next
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
	result := mergeKLists3(lists)
	for result != nil {
		fmt.Printf("sum = %+v", *result)
		result = result.Next
	}

	//fmt.Printf("sum = %+v", *l1)
}
