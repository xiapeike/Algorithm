package main

import (
	"fmt"
	"sort"
)

/*
https://leetcode.cn/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
*/

/*
1. 采用两个指针 i = m - 1, j = n - 1 倒序遍历两个数组 k = m+ n -1
2. 若发现nums1[i] < nums2[j] nums[k--] = nums[j] j--
3. 反之   nums[k--] = nums2[i] i--
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	lenth := m + n
	i, j := m-1, n-1
	k := lenth - 1
	for ; k >= 0; k-- {
		if i < 0 {
			nums1[k] = nums2[j]
			j--
			continue
		}
		if j < 0 {
			nums1[k] = nums1[i]
			i--
			continue
		}
		if nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
	}
	fmt.Printf("nums 1 = %v", nums1)
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

func main() {
	input1 := []int{1, 2, 3, 0, 0, 0}
	input2 := []int{2, 5, 6}
	merge(input1, 3, input2, 3)
}
