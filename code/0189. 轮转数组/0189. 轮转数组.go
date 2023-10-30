package main

import "fmt"

// https://leetcode.cn/problems/rotate-array/?envType=study-plan-v2&envId=top-interview-150
/*
https://leetcode.cn/problems/rotate-array/?envType=study-plan-v2&envId=top-interview-150

*/

/*
新建一个额外数组
新数组中的值 为老数组左移k个位置 即
new[i] = old[i - k<0; i - k +len(nums); i -k ]
*/
func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

/*
利用切片底层原理
*/
func rotate1(nums []int, k int) {
	copy(nums, append(nums[len(nums)-k:], nums[0:len(nums)-k]...))
	fmt.Printf("result = %v", nums)

}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	target := 3
	rotate1(input, target)
}
