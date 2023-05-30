package main

import (
	"code/util"
	"fmt"
)

// https://leetcode.cn/problems/edit-distance/
/*
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
*/

/*
(一)、当word1[i]==word2[j]时,由于遍历到了i和j,说明word1的0~i-1和word2的0~j-1的匹配结果已经生成,
由于当前两个字符相同,因此无需做任何操作,dp[i][j]=dp[i-1][j-1]
(二)、当word1[i]!=word2[j]时,可以进行的操作有3个:

	① 替换操作:可能word1的0~i-1位置与word2的0~j-1位置的字符都相同,
	     只是当前位置的字符不匹配,进行替换操作后两者变得相同,
	     所以此时dp[i][j]=dp[i-1][j-1]+1(这个加1代表执行替换操作)
	②删除操作:若此时word1的0~i-1位置与word2的0~j位置已经匹配了,
	   此时多出了word1的i位置字符,应把它删除掉,才能使此时word1的0~i(这个i是执行了删除操作后新的i)
	   和word2的0~j位置匹配,因此此时dp[i][j]=dp[i-1][j]+1(这个加1代表执行删除操作)
	③插入操作:若此时word1的0~i位置只是和word2的0~j-1位置匹配,
	    此时只需要在原来的i位置后面插入一个和word2的j位置相同的字符使得
	    此时的word1的0~i(这个i是执行了插入操作后新的i)和word2的0~j匹配得上,
	    所以此时dp[i][j]=dp[i][j-1]+1(这个加1代表执行插入操作)
	④由于题目所要求的是要最少的操作数:所以当word1[i] != word2[j] 时,
	    需要在这三个操作中选取一个最小的值赋格当前的dp[i][j]

(三)总结:状态方程为:
if(word1[i] == word2[j]):

	dp[i][j] = dp[i-1][j-1]

else:

	min(dp[i-1][j-1],dp[i-1][j],dp[i][j-1])+1
*/
func minDistance072(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, len(word1))
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 0
		}
	}

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}

	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				min1 := util.Min(dp[i-1][j-1], dp[i-1][j])
				min2 := util.Min(min1, dp[i][j-1])
				dp[i][j] = min2 + 1
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	result := minDistance072("intention", "execution")
	fmt.Printf("result = %v", result)
}
