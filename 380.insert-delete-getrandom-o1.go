/*
 * @lc app=leetcode id=380 lang=golang
 *
 * [380] Insert Delete GetRandom O(1)
 */

package main

import (
	"math/rand"
)

// @lc code=start
type RandomizedSet struct {
	index map[int]int
	vals  []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{vals: []int{}, index: make(map[int]int)}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.index[val]; ok {
		return false
	}
	rs.vals = append(rs.vals, val)
	rs.index[val] = len(rs.vals) - 1
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	if _, ok := rs.index[val]; !ok {
		return false
	}
	i := rs.index[val]
	last := rs.vals[len(rs.vals)-1]

	rs.vals[i], rs.vals[len(rs.vals)-1] = rs.vals[len(rs.vals)-1], rs.vals[i]
	rs.vals = rs.vals[:len(rs.vals)-1]

	rs.index[last] = i
	delete(rs.index, val)

	return true
}

func (rs *RandomizedSet) GetRandom() int {
	random := rand.Intn(len(rs.vals))
	return rs.vals[random]
}

/**
 * Your RandokizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
