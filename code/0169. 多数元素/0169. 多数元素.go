package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/majority-element/?envType=study-plan-v2&envId=top-interview-150
/*
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/

/*
方法一：考虑先排序，然后去中间位置的元素
*/
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

/*
方法2：
时间负责都o(n) 空间复杂度o(1)
摩尔投票法（Boyer–Moore majority vote algorithm），也被称作「多数投票法」，算法解决的问题是：如何在任意多的候选人中（选票无序），选出获得票数最多的那个。

算法可以分为两个阶段：

对抗阶段：分属两个候选人的票数进行两两对抗抵消
计数阶段：计算对抗结果中最后留下的候选人票数是否有效

根据上述的算法思想，我们遍历投票数组，将当前票数最多的候选人与其获得的（抵消后）票数分别存储在 major 与 count 中。

当我们遍历下一个选票时，判断当前 count 是否为零：

若 count == 0，代表当前 major 空缺，直接将当前候选人赋值给 major，并令 count++
若 count != 0，代表当前 major 的票数未被完全抵消，因此令 count--，即使用当前候选人的票数抵消 major 的票数
*/
func majorityElement2(nums []int) int {
	major := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if major == num {
			count += 1
		} else {
			count -= 1
		}
	}

	return major
}

func main() {
	input := []int{3, 2, 3}
	result := majorityElement2(input)
	fmt.Printf("result = %v", result)
}
