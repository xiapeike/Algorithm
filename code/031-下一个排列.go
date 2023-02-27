package main

import "fmt"

//https://leetcode-cn.com/problems/next-permutation/
/*
整数数组的一个 排列 就是将其所有成员以序列或线性顺序排列。

例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那
么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，
那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。
*/

func reverse_arr(arr []int, start int, end int) {
	for start < end {
		// 交换
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}
}

/*
1、从后向前遍历，找到第一个a[i]<a[j]  [..., i, j, ...],此时有[j, end]是降序
2、从后向前遍历，找到一个k，使a[k] > a[i], 并进行交换。a[k + 1]> a[k] > a[i] > a[k - 1]
3、此时交换后[j, end]任是降序，需要逆置，找到最小排列
4、如果在步骤 1 找不到符合的相邻元素对，说明当前 [begin,end) 为一个降序顺序，则直接跳到步骤 3
*/
func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}
	// 从后向前遍历，找到第一个a[i]<a[j]
	i := len(nums) - 2
wc:
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			// 从后向前遍历，找到一个k，使a[k] > a[i], 并进行交换。
			for k := len(nums) - 1; k > i; k-- {
				if nums[k] > nums[i] {
					// 交换
					temp := nums[k]
					nums[k] = nums[i]
					nums[i] = temp
					break wc
				}
			}
		}
	}
	reverse_arr(nums, i+1, len(nums)-1)
}

func main() {
	input := []int{2, 3, 1}
	nextPermutation(input)
	fmt.Printf("target = %v", input)
}
