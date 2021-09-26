package main

import (
	"fmt"
)

// 方法一：两次循环遍历
func twoSum1(nums []int, target int) []int {
	var returnList [2]int
Loop:
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				returnList[0] = i
				returnList[1] = j
				break Loop
			}
		}
	}
	return returnList[:]
}

// 方法二：构建一个map表，用空间换时间
func twoSum2(nums []int, target int) []int {
	var returnList [2]int
	mapList := make(map[int]int)

	// 遍历一次数组，将每个值作为key插入到map中，方便再次查询
	for i := 0; i < len(nums); i++ {
		mapList[nums[i]] = i
	}
Loop:
	for i := 0; i < len(nums); i++ {
		j, ok := mapList[target-nums[i]]
		if ok == true && i != j {
			returnList[0] = i
			returnList[1] = j
			break Loop
		}
	}
	return returnList[:]
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	sum := twoSum2(nums, target)
	fmt.Printf("sum = %v", sum)
}
