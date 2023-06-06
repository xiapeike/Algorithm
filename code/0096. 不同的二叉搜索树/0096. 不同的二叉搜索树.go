package main

import (
	"fmt"
)

// https://leetcode.cn/problems/unique-binary-search-trees/
/*
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
*/

/*
递归  超时
1. 遍历i 属于1-n， 以i为根， [1, i-1] 为左树结点， [i+1, n] 为右树结点
2. 分别计算左树、右数结点构成二叉搜索树的种树， 并进行相乘返回
3. 递归退出条件：
	左（右）子树为空 或者为单结点，数量返回1
*/

func numTrees096(n int) int {
	result := 0
	nums := make([]int, n)

	var dps func([]int) int
	dps = func(nums []int) int {
		if len(nums) <= 1 {
			return 1
		}
		sum := 0

		for i := 0; i < len(nums); i++ {
			// 计算左子树种树
			left_num := dps(nums[0:i])

			// 计算右子树种树
			right_num := dps(nums[i+1 : len(nums)])

			sum = sum + left_num*right_num
		}

		return sum
	}
	result = dps(nums)
	return result

}

/*
对于第一种方法，我们采用空间换时间
对于每一个子树计算结果，我们记录下来，方便后续使用
*/
func numTrees0962(n int) int {
	result := 0
	memo := make([]int, n+1)
	var dps func(int) int
	dps = func(n int) int {
		if n == 1 || n == 0 {
			return 1
		}
		if memo[n] > 0 {
			return memo[n]
		}
		count := 0
		for i := 0; i <= n-1; i++ {
			count += dps(i) * dps(n-i-1)
		}
		memo[n] = count
		return count
	}
	result = dps(n)
	return result

}

/*
动态规划
https://leetcode.cn/problems/unique-binary-search-trees/solutions/329807/bu-tong-de-er-cha-sou-suo-shu-by-leetcode-solution/
1. G(n) 代表长度为 n 的序列能构成的不同二叉搜索树的个数
2. F(i,n): 以 i 为根、序列长度为 n 的不同二叉搜索树个数 (1≤i≤n)
3. 可得出  G(n) = F(1,n)+F(2,n)+ ... F(n,n)
4. 又 F(i,n) = G(i-1)*G(n-i)   注：以i为根的树，总数为左子树的个数*右子树的个数
5. G(0) = 1  G(1) = 1
*/
func numTrees0963(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]

}

func main() {
	result := numTrees0963(4)
	fmt.Printf("result = %v", result)
}
