package main

// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
import (
	"fmt"
	"strconv"
)

// 方法一：分别遍历两个数组
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result float64 = 0

	// 合并数组
	var mergeArray []int
	var i, j int = 0, 0
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	for i < len(nums1) {
		val1 := nums1[i]
		if j == len(nums2) {
			mergeArray = append(mergeArray, val1)
			i++
			continue
		}
		for j < len(nums2) {
			val2 := nums2[j]
			if val1 <= val2 {
				mergeArray = append(mergeArray, val1)
				i++
				break
			} else {
				mergeArray = append(mergeArray, val2)
				j++
			}
		}
	}

	if j < len(nums2) {
		for j < len(nums2) {
			val2 := nums2[j]
			mergeArray = append(mergeArray, val2)
			j++
		}
	}
	lenth := len(mergeArray)
	if lenth > 0 && lenth%2 == 1 {
		result = (float64)(mergeArray[(lenth-1)/2])
	} else if lenth > 0 {
		result = ((float64)(mergeArray[(lenth-1)/2]) + (float64)(mergeArray[lenth/2])) / 2
	}
	result, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", result), 64)
	return result
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	target := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("sum = %5f", target)
}
