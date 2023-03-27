package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/multiply-strings/
/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。
*/

/*
"123" * "456"
"123" * "6" + "123" * "5" * "10" + "123" * "4" * "100"
*/
func multiply043(num1 string, num2 string) string {
	result := ""
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	var dfs func(string, int, int)

	dfs = func(num1 string, num2_index int, zero_cn int) {
		if num2_index < 0 {
			return
		}

		b := int(num2[num2_index] - '0')
		multiply_num := multiply2(num1, b)
		for i := 0; i < zero_cn; i++ {
			multiply_num = multiply_num + "0"
		}
		result = string_sum(result, multiply_num)
		zero_cn = zero_cn + 1
		num2_index = num2_index - 1
		dfs(num1, num2_index, zero_cn)
	}
	dfs(num1, len(num2)-1, 0)
	return result
}

func multiply2(num1 string, b int) string {
	result := ""
	jinwei := 0
	if b == 0 {
		return "0"
	}
	for i := len(num1) - 1; i >= 0; i-- {
		a := int(num1[i] - '0')
		mul := a*b + jinwei
		ge := mul % 10
		shi := mul / 10
		jinwei = shi
		result = strconv.Itoa(ge) + result
	}
	if jinwei != 0 {
		result = strconv.Itoa(jinwei) + result
	}
	return result
}

func string_sum(str1 string, str2 string) string {
	result := ""
	jinwei := 0
	// str1一定比str2长
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}
	len1 := len(str1) - 1
	len2 := len(str2) - 1

	for i := 0; i <= len2; i++ {
		a := int(str1[len1-i] - '0')
		b := int(str2[len2-i] - '0')
		sum := a + b + jinwei
		if sum >= 10 {
			jinwei = 1
			sum = sum - 10
		} else {
			jinwei = 0
		}
		result = strconv.Itoa(sum) + result
	}

	for i := len2 + 1; i <= len1; i++ {
		a := int(str1[len1-i] - '0')
		sum := a + jinwei
		if sum >= 10 {
			jinwei = 1
			sum = sum - 10
		} else {
			jinwei = 0
		}
		result = strconv.Itoa(sum) + result
	}
	if jinwei != 0 {
		result = strconv.Itoa(jinwei) + result
	}
	return result
}
func main() {
	input := "123456789"
	input2 := "987654321"
	result := multiply043(input, input2)
	fmt.Printf("result = %v", result)
}
