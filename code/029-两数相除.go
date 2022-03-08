package main

import (
	"fmt"
	"math"
)

//https://leetcode-cn.com/problems/divide-two-integers/
/*
给定两个整数，被除数dividend和除数divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数dividend除以除数divisor得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2

*/

// 不断用被除数减除数，直到被除数小于除数退出循环
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 { // 考虑被除数为最小值的情况
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 { // 考虑除数为最小值的情况
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if dividend == 0 { // 考虑被除数为 0 的情况
		return 0
	}

	flag := 1
	if dividend < 0 {
		flag = flag * -1
		dividend = dividend * -1
	}

	if divisor < 0 {
		flag = flag * -1
		divisor = divisor * -1
	}

	results := 0
	for dividend >= divisor {
		results = results + 1
		dividend = dividend - divisor
	}

	return results * flag
}

// 方法一极端情况处理时间过长，可使用二分法加速退出循环的过程
func divide2(dividend int, divisor int) int {
	if dividend == math.MinInt32 { // 考虑被除数为最小值的情况
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 { // 考虑除数为最小值的情况
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if dividend == 0 { // 考虑被除数为 0 的情况
		return 0
	}

	flag := 1
	if dividend < 0 {
		flag = flag * -1
		dividend = dividend * -1
	}

	if divisor < 0 {
		flag = flag * -1
		divisor = divisor * -1
	}
	results := 0
	if dividend >= divisor {
		results = div(dividend, divisor, 1)
	}

	return results * flag
}

func div(dividend int, divisor int, results int) int {
	divisor_temp := divisor
	if dividend < divisor {
		return 0
	}
	for dividend >= divisor_temp {
		if divisor_temp*2 <= dividend {
			results = results * 2
			divisor_temp = divisor_temp * 2
		} else {
			dividend = dividend - divisor_temp
			results = results + div(dividend, divisor, 1)
		}

	}
	return results
}

func main() {
	target := divide2(11, 3)
	fmt.Printf("sum = %v", target)

}
