/*
 * @lc app=leetcode id=341 lang=golang
 *
 * [341] Flatten Nested List Iterator
 */
package main

type NestedInteger struct{}

func (ni NestedInteger) IsInteger() bool {
	return false
}

func (ni NestedInteger) GetInteger() int {
	return 0
}
func (n *NestedInteger) SetInteger(value int)    {}
func (ni *NestedInteger) Add(elem NestedInteger) {}
func (ni NestedInteger) GetList() []*NestedInteger {
	return []*NestedInteger{}
}

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	flattened []int
	index     int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		flattened: flatten(nestedList),
		index:     0,
	}
}

func flatten(nested []*NestedInteger) []int {
	var result []int
	for _, ni := range nested {
		if ni.IsInteger() {
			result = append(result, ni.GetInteger())
		} else {
			result = append(result, flatten(ni.GetList())...)
		}
	}
	return result
}

func (ni *NestedIterator) Next() int {
	val := ni.flattened[ni.index]
	ni.index++
	return val
}

func (ni *NestedIterator) HasNext() bool {
	return ni.index < len(ni.flattened)
}

// @lc code=end
