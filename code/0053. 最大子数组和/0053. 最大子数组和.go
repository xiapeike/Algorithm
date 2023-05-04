package main

import (
	"fmt"
)

// https://leetcode.cn/problems/powx-n/
/*
实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，x^n ）。
*/

/*
此方案n次循环会超时
*/
func myPow050(x float64, n int) float64 {
	result := 1.00000
	if n == 0 {
		return result
	} else if n > 0 {
		for i := 0; i < n; i++ {
			result = result * x
		}
	} else {
		n = n * -1
		for i := 0; i < n; i++ {
			result = result * x
		}
		result = 1 / result
		if x < 0 && n%2 == 1 {
			result = -1 * result
		}
	}

	return result
}

/*
快速幂法 时间复杂度降低至 O(logn)
快速幂 + 迭代
*/
func myPow0502(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}
func quickMul(x float64, N int) float64 {
	ans := 1.0
	// 贡献的初始值为 x
	x_contribute := x
	// 在对 N 进行二进制拆分的同时计算答案
	for N > 0 {
		if N%2 == 1 {
			// 如果 N 二进制表示的最低位为 1，那么需要计入贡献
			ans *= x_contribute
		}
		// 将贡献不断地平方
		x_contribute *= x_contribute
		// 舍弃 N 二进制表示的最低位，这样我们每次只要判断最低位即可
		N /= 2
	}
	return ans
}

func main() {
	result := myPow0502(-13.62608, 3)
	fmt.Printf("result = %v", result)
}
