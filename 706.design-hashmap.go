/*
 * @lc app=leetcode id=706 lang=golang
 *
 * [706] Design HashMap
 */

package main

import "fmt"

// @lc code=start
type Node706 struct {
	key   int
	value int
	next  *Node706
}

func (n *Node706) Add(key int, value int) *Node706 {
	return &Node706{key: key, value: value, next: n}
}

func (n *Node706) Remove(key int) *Node706 {
	if n == nil {
		return nil
	}
	if n.key == key {
		return n.next
	}
	n.next = n.next.Remove(key)
	return n
}

func (n *Node706) Get(key int) *Node706 {
	for current := n; current != nil; current = current.next {
		if current.key == key {
			return current
		}
	}
	return nil
}

type MyHashMap struct {
	list  []*Node706
	size  int
	items int
}

func Constructor() MyHashMap {
	return constructor(1)
}

func constructor(size int) MyHashMap {
	if size < 1 {
		size = 1
	}
	return MyHashMap{size: size, list: make([]*Node706, size)}
}

func (this *MyHashMap) put(key int, value int) {
	list := this.list[this.hash(key)]
	if node := list.Get(key); node != nil {
		node.value = value
	} else {
		list = list.Add(key, value)
		this.list[this.hash(key)] = list
		this.items++
	}
}

func (this *MyHashMap) Put(key int, value int) {
	this.put(key, value)
	this.resize()
}

func (this *MyHashMap) Get(key int) int {
	list := this.list[this.hash(key)]
	if list.Get(key) == nil {
		return -1
	} else {
		return list.Get(key).value
	}
}

func (this *MyHashMap) remove(key int) {
	list := this.list[this.hash(key)]
	list = list.Remove(key)
	this.list[this.hash(key)] = list
	this.items--
}

func (this *MyHashMap) Remove(key int) {
	this.remove(key)
	this.resize()
}

func (this *MyHashMap) resize() {
	var newSize int
	if this.items*4 < this.size {
		newSize = this.size / 2
	} else if this.items*4 > this.size*3 {
		newSize = this.size * 2
	} else {
		return
	}

	newList := constructor(newSize)
	for _, list := range this.list {
		for current := list; current != nil; current = current.next {
			newList.put(current.key, current.value)
		}
	}
	this.list = newList.list
	this.size = newList.size
	this.items = newList.items
}

func (this *MyHashMap) hash(key int) int {
	key = ((key >> 16) ^ key) * 0x45d9f3b
	return key % this.size
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
// @lc code=end
func main() {
	obj := Constructor()
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1) == 1)
	obj.Get(1)
	fmt.Println(obj.Get(3) == -1)
	obj.Put(2, 1)
	fmt.Println(obj.Get(2) == 1)
	obj.Remove(2)
	fmt.Println(obj.Get(2) == -1)
}
