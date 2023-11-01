package main

import "fmt"

/*
https://leetcode.cn/problems/candy/?envType=study-plan-v2&envId=top-interview-150
n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。

你需要按照以下要求，给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
*/

/*
遍历数组，为每个孩子构建好入度矩阵和度数矩阵
遍历度数为0的结点，分配糖果数（初始为1，后续每轮循环加一），并将该结点从其他结点的入度矩阵中移除，

该方法会超时
*/
func candy(ratings []int) int {
	n := len(ratings) // 结点数
	// 构建入度矩阵和度数矩阵
	degree := make([][]int, n)
	degree_num := make([]int, n) // 度数矩阵，用于表示当前结点的最新度数，-1表示已经入队过
	// 初始话入度矩阵
	for i := 0; i < n; i++ {
		degree[i] = make([]int, n)
		for j := 0; j < n; j++ {
			degree[i][j] = 0
		}
	}
	// 初始话度数矩阵
	for i := 0; i < n-1; i++ {
		if ratings[i] < ratings[i+1] {
			// i + 1 结点的入度数加1，入度矩阵加一
			degree_num[i+1]++
			degree[i+1][i]++
		} else if ratings[i] > ratings[i+1] {
			// i 结点的入度数加1，入度矩阵加一
			degree_num[i]++
			degree[i][i+1]++
		}
	}

	// 队列，用于存放度为0的结点，队不为空，说明依然有结点需要遍历
	queue := make([]int, n)
	top := -1
	// 将度数为0的结点入队
	for i := 0; i < n; i++ {
		if degree_num[i] == 0 {
			top++
			queue[top] = i
		}
	}

	// 初始化糖果数
	num := 0
	result := 0
	for top >= 0 {
		num++
		for top >= 0 {
			// 出队
			k := queue[top]
			top--
			result += num
			degree_num[k]--
			// 从所有结点的入度矩阵中移除k
			// 此处会超时，需要剪枝，只考虑相邻结点
			//for i := 0; i < n; i++ {
			//	if degree[i][k] > 0 {
			//		degree[i][k] = 0
			//		degree_num[i]--
			//	}
			//}
			if k > 0 && degree[k-1][k] > 0 {
				degree[k-1][k] = 0
				degree_num[k-1]--
			}

			if k < n-1 && degree[k+1][k] > 0 {
				degree[k+1][k] = 0
				degree_num[k+1]--
			}

		}

		// 将新的结点入队
		for i := 0; i < n; i++ {
			if degree_num[i] == 0 {
				top++
				queue[top] = i
			}
		}
	}
	return result

}

func candy2(ratings []int) int {
	// 分发糖果， 贪心算法
	need := make([]int, len(ratings))
	sum := 0
	// 保证条件1，每人先分一颗糖
	for i := range need {
		need[i] = 1
	}
	// 满足条件2，分高的比相邻的多， 左中右，所以分向后遍历及向前遍历两种
	// 先由前向后遍历，遇到右边大于左边的rating的，糖果+1 -->>
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i] < ratings[i+1] {
			need[i+1] = need[i] + 1
		}
	}
	// 由后向前遍历，遇到左边大于右边的rating，此时应该考虑左边的应取(need[i-1], need[i]+1)的大值，
	// 如果都已经满足左边分高，且糖果多，就不用另外分配糖果了，节省糖果
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			need[i-1] = func(a, b int) int {
				if a > b {
					return a
				}
				return b
			}(need[i-1], need[i]+1)
		}
	}
	// 最后统计糖果的总数
	for _, v := range need {
		sum += v
	}
	return sum
}

func main() {
	input := []int{1, 0, 2}
	result := candy2(input)
	fmt.Printf("result = %v", result)
}
