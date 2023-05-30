package main

import (
	"code/util"
	"fmt"
)

// https://leetcode.cn/problems/largest-rectangle-in-histogram/
/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。
*/

/*
方法一：暴力解法（超时）
可以枚举以每个柱形为高度的最大矩形的面积。

具体来说是：依次遍历柱形的高度，对于每一个高度分别向两边扩散，求出以当前高度为矩形的最大宽度多少。
*/

func largestRectangleArea084(heights []int) int {
	result := 0

	for i := 0; i < len(heights); i++ {
		max := 0
		// 遍历每一根柱子
		// 往左延伸
		for l := i; l >= 0; l-- {
			if heights[l] >= heights[i] {
				max = max + heights[i]
			} else {
				break
			}
		}
		// 往右延伸
		for r := i; r < len(heights); r++ {
			if heights[r] >= heights[i] {
				max = max + heights[i]
			} else {
				break
			}
		}

		max = max - heights[i]

		if max > result {
			result = max
		}

	}

	return result
}

/*
方法二：空间换时间  采用单调栈（栈内元素顺序是单调的）
这题考的基础模型其实就是：在一维数组中对每一个数找到第一个比自己小的元素。这类“在一维数组中找第一个满足某种条件的数”的场景就是典型的单调栈应用场景。

*/

func largestRectangleArea0842(heights []int) int {
	//单调栈（单调递增） 从栈顶到栈底依次增加
	stack := make([]int, 0)
	stack = append(stack, -1)    //stack的哨兵，方便确定左边界
	heights = append(heights, 0) //添加一个哨兵，减少代码量
	ln := len(heights)
	res := 0 //结果

	for i := 0; i < ln; i++ {
		//因为我们无法访问heights[-1]，所以限制len(stack) > 1
		for len(stack) > 1 && heights[stack[len(stack)-1]] > heights[i] {
			//栈顶元素，也就是当前要求的矩形柱子的下标
			top := stack[len(stack)-1]
			//出栈
			stack = stack[:len(stack)-1]
			//左边界（栈顶元素的后一个元素）
			l := stack[len(stack)-1]
			//矩形面积：(右边界-左边界-1) * 高度
			//右边界就是i
			//高度就是以栈顶元素为下标的柱子的高度
			//左边界就是栈顶元素的下一位元素（因为我们添加了哨兵-1，所以这公式依旧成立）
			res = util.Max(res, (i-l-1)*heights[top])
		}
		stack = append(stack, i)

	}

	return res
}

func main() {
	input := []int{2, 1, 5, 6, 2, 3}
	//input := []int{2, 4}
	//input1 := "AB"
	result := largestRectangleArea0842(input)
	fmt.Printf("result = %v", result)
}
