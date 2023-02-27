package util

import (
	"fmt"
	"sort"
)

//	func main() {
//		lockTest()
//	}
func lockTest() {
	var mp map[string]int
	mp = make(map[string]int)
	mp["1"] = 1
	mp["5"] = 45
	mp["6"] = 3
	mp["4"] = 22
	mp["2"] = 5
	mp["3"] = 16

	sortMap(mp)

	sortMap2(mp)
}

/*
*
根据key排序
*/
func sortMap2(mp map[string]int) {
	var newMp = make([]string, 0)
	for k, _ := range mp {
		newMp = append(newMp, k)
	}
	sort.Strings(newMp)
	for _, v := range newMp {
		fmt.Println("根据key排序后的新集合》》   key:", v, "    value:", mp[v])
	}
}

/*
*
根据value排序
*/
func sortMap(mp map[string]int) {
	reverseMap := make(map[int][]string)

	for key, value := range mp {
		reverseMap[value] = append(reverseMap[value], key)
	}
	fmt.Println("reverseMap:", reverseMap)
}
