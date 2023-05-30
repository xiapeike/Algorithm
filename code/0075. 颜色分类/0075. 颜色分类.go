package main

import (
	"fmt"
)

// https://leetcode.cn/problems/sort-colors/
/*
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

必须在不使用库内置的 sort 函数的情况下解决这个问题。
*/

/*
冒泡排序
*/
func sortColors075(nums []int) {
	for i := len(nums) - 1; i >= 1; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	fmt.Printf("result = %v", nums)
}

/*
遍历一趟，记录0，1，2的个数
再遍历一趟赋值
*/
func sortColors0752(nums []int) {
	count_0, count_1, count_2 := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count_0++
		}
		if nums[i] == 1 {
			count_1++
		}
		if nums[i] == 2 {
			count_2++
		}
	}

	for i := 0; i < len(nums); i++ {
		if i < count_0 {
			nums[i] = 0
		} else if i < count_0+count_1 {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
	fmt.Printf("result = %v", nums)
}

/*
双指针法
https://leetcode.cn/problems/sort-colors/solutions/437968/yan-se-fen-lei-by-leetcode-solution/
*/
func sortColors0753(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		for ; i <= p2 && nums[i] == 2; p2-- {
			nums[i], nums[p2] = nums[p2], nums[i]
		}
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}

func main() {
	input := []int{2, 0, 2, 1, 1, 0}
	sortColors0752(input)
	//fmt.Printf("result = %v", result)
}
