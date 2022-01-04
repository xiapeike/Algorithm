package main

// https://leetcode-cn.com/problems/container-with-most-water/
/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点(i,ai) 。在坐标内画 n 条垂直线，垂直线 i的两个端点分别为(i,ai) 和 (i, 0) 。
找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
*/

import (
	"fmt"
)

func getMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*解法一：循环遍历，会超时*/
func maxArea(height []int) int {
	max_area := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			w := j - i
			h := getMin(height[i], height[j])
			if max_area < w*h {
				max_area = w * h
			}
		}
	}
	return max_area
}

/*解法一：双指针*/
func maxArea2(height []int) int {
	max_area := 0
	i, j := 0, len(height)-1
	for i < j {
		h := getMin(height[i], height[j])
		w := j - i
		if max_area < w*h {
			max_area = w * h
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return max_area
}
func main() {
	input := []int{4, 3, 2, 1, 4}
	//target := longestPalindrome(str)
	result := maxArea2(input)
	fmt.Printf("result = %v", result)
}
