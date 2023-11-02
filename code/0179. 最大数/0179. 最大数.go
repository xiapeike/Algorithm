package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

/*
https://leetcode.cn/problems/largest-number/
给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
*/

/*
依次比较相邻两个数a和b
如若ab < ba 则调换a 和b的顺序
冒泡排序
*/
func largestNumber(nums []int) string {
	var result bytes.Buffer

	// 冒泡排序
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			a, b := nums[j], nums[j+1]
			if multiply(a, b) {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	if nums[0] == 0 {
		return "0"
	}
	for i := 0; i < len(nums); i++ {
		result.WriteString(strconv.Itoa(nums[i]))
	}
	return result.String()
}

func multiply(a, b int) bool {
	ab, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	ba, _ := strconv.Atoi(strconv.Itoa(b) + strconv.Itoa(a))
	if ab < ba {
		return true
	}
	return false
}

/*
标准做法
*/
func largestNumber2(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sy*x+y > sx*y+x
	})
	if nums[0] == 0 {
		return "0"
	}
	ans := []byte{}
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}

func main() {
	input := []int{3, 30, 34, 5, 9}
	result := largestNumber(input)
	fmt.Printf("result = %v", result)
}
