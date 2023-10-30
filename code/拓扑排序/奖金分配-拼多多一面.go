package main

import "fmt"

/*
【题目描述】 https://blog.csdn.net/weixin_51995229/article/details/126998658?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_baidulandingword~default-4-126998658-blog-104024552.235^v38^pc_relevant_default_base3&spm=1001.2101.3001.4242.3&utm_relevant_index=7
由于无敌的凡凡在2005年世界英俊帅气男总决选中胜出，Yali Company总经理Mr.Z心情好，决定给每位员工发奖金。公司决定以每个人本年在公司的贡献为标准来计算他们得到奖金的多少。

于是Mr.Z下令召开m方会谈。每位参加会谈的代表提出了自己的意见：“我认为员工a的奖金应该比b高！”Mr.Z决定要找出一种奖金方案，满足各位代表的意见，且同时使得总奖金数最少。每位员工奖金最少为100元。

【输入】
第一行两个整数n,m，表示员工总数和代表数；

以下m行，每行2个整数a,b，表示某个代表认为第a号员工奖金应该比第b号员工高。

【输出】
若无法找到合理方案，则输出“Poor Xed”；否则输出一个数表示最少总奖金。

【输入样例】
2 1
1 2
【输出样例】
201

【输入样例】
4 4
1 2
1 3
1 4
2 3
【输出样例】
403
【提示】
【数据规模】

80％的数据满足：n≤1000，m≤2000；

100％的数据满足：n≤10000，m≤20000。

*/

/*
思路：
一、构建每个员工的入度矩阵
如 1号员工 > 2号员工工资   则 1的入度矩阵中，添加2号员工

二、遍历每个员工的入度矩阵，如果入度为0， 奖金+100， 并将该员工从入度矩阵中移除，且所有入度矩阵都需移除该结点
结束该轮遍历后，下一轮奖金+1，直到所以员工均被移除，结束遍历

难点：
入度矩阵数据结构如何定义？
degree []int  存储每个员工的入度数
1：[0,1，1，1]  入度二维矩阵，标识每个员工所有的具体入度员工
*/
func solution(n int, m int, array [][]int) int {
	result := 0
	// 构建degree矩阵
	degree := make([]int, n)
	degree_map := make([][]int, n, n)
	for i := 0; i < n; i++ {
		degree_map[i] = make([]int, n)
		for j := 0; j < n; j++ {
			degree_map[i][j] = 0
		}
	}
	for i := 0; i < m; i++ {
		degree[array[i][0]-1]++
		degree_map[array[i][0]-1][array[i][1]-1] = 1

	}

	// 循环遍历degree矩阵
	gold := 100
	for {
		count := 0
		for i := 0; i < len(degree); i++ {
			if degree[i] == 0 {
				result += gold
				degree[i] = -1 // 标记移除
			}
		}
		for i := 0; i < len(degree); i++ {
			if degree[i] == -1 {
				// 从其他所有员工的度矩阵中移除
				for j := 0; j < len(degree_map); j++ {
					if i != j && degree_map[j][i] == 1 {
						degree[j]--
						degree_map[j][i] = 0
					}
				}
				count++
			}
		}
		if count == len(degree) {
			break
		}
		gold++
	}

	return result
}

func main() {
	// case 1
	//n, m := 4, 4
	//array := [][]int{
	//	{1, 2},
	//	{1, 3},
	//	{1, 4},
	//	{2, 3},
	//}

	// case 2
	n, m := 2, 1
	array := [][]int{
		{1, 2},
	}

	result := solution(n, m, array)
	fmt.Printf("result = %v", result)
}
