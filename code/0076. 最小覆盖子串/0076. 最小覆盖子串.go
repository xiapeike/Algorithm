package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/minimum-window-substring/
/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。



注意：

对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。
*/

/*
暴力法, 会超时
*/
func minWindow076(s string, t string) string {
	m, n := len(s), len(t)
	if m < n {
		return ""
	}
	// 窗口长度从n开始递增，遍历s，判断每个窗口内，是否包含t
	for len_window := n; len_window <= m; len_window++ {
		// 根据该窗口长度，遍历字符串s
		for i := 0; i < m-len_window+1; i++ {
			visit_map := make(map[int]bool)
			// 初始化visit_map
			for j := i; j < i+len_window; j++ {
				visit_map[j] = false
			}

			// 判断该窗口字符串是否包含t的每个字符，如果包含，就返回该窗口字符串
			flag_array := make([]bool, len(t))
			for j := 0; j < len(t); j++ {
				flag_array[j] = false
			}
			for k := 0; k < len(t); k++ {
				for l := i; l < i+len_window; l++ {
					v, _ := visit_map[l]
					if t[k] == s[l] && v == false {
						flag_array[k] = true
						visit_map[l] = true
						break
					}
				}
			}

			flag_true_num := 0
			for j := 0; j < len(t); j++ {
				if flag_array[j] == true {
					flag_true_num++
				}
			}

			if flag_true_num == len(t) {
				return string(s[i : i+len_window])
			} else if flag_true_num == 0 {
				i = i + len_window - 1
			}
		}
	}
	return ""
}

/*
滑动窗口法
。在滑动窗口类型的问题中都会有两个指针，一个用于「延伸」现有窗口的 r 指针，和一个用于「收缩」窗口的 l 指针。在任意时刻，只有一个指针运动，而另一个保持静止。我们在 s 上滑动窗口，通过移动rr 指针不断扩张窗口。当窗口包含 t 全部所需的字符后，如果能收缩，我们就收缩窗口直到得到最小窗口。
*/
func minWindow0762(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}

func main() {
	input := "ADOBECODEBANC"
	//input1 := "AB"
	result := minWindow0762(input, "ABC")
	fmt.Printf("result = %v", result)
}
