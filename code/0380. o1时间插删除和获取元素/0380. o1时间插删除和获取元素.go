package main

import (
	"fmt"
	"math/rand"
)

/*
https://leetcode.cn/problems/insert-delete-getrandom-o1/?envType=study-plan-v2&envId=top-interview-150

实现RandomizedSet 类：

RandomizedSet() 初始化 RandomizedSet 对象
bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。
*/

/*
变长数组 + 哈希表
*/
type RandomizedSet struct {
	nums    []int       // 用于o(1)操作
	indices map[int]int // 用于判断元素是否存在， key为元素， val为key在nums数组中的下标，删除时用到
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.indices[val]; ok {
		return false
	}
	rs.indices[val] = len(rs.nums)
	rs.nums = append(rs.nums, val)
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	id, ok := rs.indices[val]
	if !ok {
		return false
	}
	last := len(rs.nums) - 1
	rs.nums[id] = rs.nums[last]
	rs.indices[rs.nums[id]] = id
	rs.nums = rs.nums[:last]
	delete(rs.indices, val)
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.nums[rand.Intn(len(rs.nums))]
}

func main() {
	//Your RandomizedSet object will be instantiated and called as such:
	obj := Constructor()
	param_1 := obj.Insert(1)
	param_2 := obj.Remove(2)
	param_3 := obj.Remove(1)
	param_4 := obj.GetRandom()
	fmt.Printf("param_1 = %v\n", param_1)
	fmt.Printf("param_2 = %v\n", param_2)
	fmt.Printf("param_3 = %v\n", param_3)
	fmt.Printf("param_3 = %v\n", param_4)
}
