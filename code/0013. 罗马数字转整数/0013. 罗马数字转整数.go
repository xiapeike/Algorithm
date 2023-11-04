package main

import "fmt"

// https://leetcode.cn/problems/roman-to-integer/?envType=study-plan-v2&envId=top-interview-150
/*罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1 。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。
*/

var (
	map_roma = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
)

func romanToInt(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		item := map_roma[s[i]]
		if i != len(s)-1 {
			switch s[i] {
			case 'I':
				if s[i+1] == 'V' || s[i+1] == 'X' {
					item = map_roma[s[i+1]] - map_roma[s[i]]
					i++
				}

			case 'X':
				if s[i+1] == 'L' || s[i+1] == 'C' {
					item = map_roma[s[i+1]] - map_roma[s[i]]
					i++
				}
			case 'C':
				if s[i+1] == 'D' || s[i+1] == 'M' {
					item = map_roma[s[i+1]] - map_roma[s[i]]
					i++
				}

			}
		}

		result += item
	}
	return result
}

var symbolsValue = map[string]int{
	"M":  1000,
	"CM": 900,
	"D":  500,
	"CD": 400,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

func romanToInt2(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		str := string(s[i])
		str2 := ""
		if i+1 < len(s) {
			str2 = string(s[i+1])
		}
		str3 := str + str2
		num, ok := symbolsValue[str3]
		if ok {
			result = result + num
			i = i + 1
		} else {
			result = result + symbolsValue[str]
		}
	}
	return result
}
func main() {
	input := "III"
	result := romanToInt2(input)
	fmt.Printf("result = %v", result)
}
