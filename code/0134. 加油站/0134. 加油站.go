package main

import "fmt"

/*
https://leetcode.cn/problems/gas-station/description/?envType=study-plan-v2&envId=top-interview-150
在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。

你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。

给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。
*/

/*
暴力法
分别以i号汽油站为起点，循环一圈，判断是否可行
*/
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)

	for i := 0; i < n; i++ {
		// 分别以i号汽油站为起点
		j := i
		k := 0        //
		cost_oil := 0 // 剩余油量
		for ; k < n; k++ {
			// 循环一圈
			j = j % n
			cost_oil = cost_oil + gas[j] - cost[j]
			if cost_oil < 0 {
				break
			}
			j++
		}
		if k == n && cost_oil >= 0 {
			return i
		}

	}
	return -1
}

/*
基于方法一做剪枝
分别以i号汽油站为起点，循环一圈，判断是否可行，
若不可行，直接跳过不可行的结点
*/
func canCompleteCircuit2(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		sumOfGas, sumOfCost, cnt := 0, 0, 0
		for cnt < n {
			j := (i + cnt) % n
			sumOfGas += gas[j]
			sumOfCost += cost[j]
			if sumOfCost > sumOfGas {
				break
			}
			cnt++
		}
		if cnt == n {
			return i
		} else {
			// 和方法一的关键区别所在
			i += cnt + 1
		}
	}
	return -1
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	result := canCompleteCircuit2(gas, cost)
	fmt.Printf("result = %v", result)
}
