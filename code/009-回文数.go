package main

// https://leetcode-cn.com/problems/palindrome-number/
/*
给你一个整数x，如果x是一个回文整数，返回true；否则，返回 false。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而123 不是。
*/
import (
	"fmt"
	"strconv"
)

/*
	解法一：将整数转化为字符串
*/
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)

	if len(str) == 0 {
		return false
	}

	if len(str) == 1 {
		return true
	}

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

/*
	解法二：将整数进行倒序，然后判断是否相等
*/
func isPalindromeV2(x int) bool {
	if x < 0 {
		return false
	}
	result := 0
	temp := x
	for x > 0 {
		result = 10*result + x%10
		x = x / 10
	}
	if result != temp {
		return false
	}
	return true
}

func main() {
	input := -101
	//target := longestPalindrome(str)
	result := isPalindromeV2(input)
	fmt.Printf("result = %v", result)
}
