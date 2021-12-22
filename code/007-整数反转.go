package main

// https://leetcode-cn.com/problems/reverse-integer/
/*
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围[−2^31, 2^31− 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。

*/
import (
	"fmt"
	"math"
)

/*
	解法一：使用数学的方法翻转，由于不能存储64位数，所以需要考虑边界值的计算
*/
func reverse(x int) int {
	rev := 0
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return rev

}

func main() {
	num := 123
	//target := longestPalindrome(str)
	result := reverse(num)
	fmt.Printf("result = %v", result)
}
