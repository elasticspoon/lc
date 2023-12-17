/*
 * @lc app=leetcode id=2353 lang=golang
 *
 * [2353] Design a Food Rating System
 */

package main

import (
	"container/heap"
)

// @lc code=start
type FoodRatings struct {
	Cuisines map[string]*FoodHeap
	Foods    map[string]*FoodItem
}

type FoodItem struct {
	name    string
	cuisine string
	rating  int
	index   int
}

type FoodHeap []*FoodItem

func (fh FoodHeap) Len() int { return len(fh) }

func (fh FoodHeap) Less(i, j int) bool {
	if fh[i].rating == fh[j].rating {
		return fh[i].name < fh[j].name
	}
	return fh[i].rating > fh[j].rating
}

func (fh FoodHeap) Swap(i, j int) {
	fh[i], fh[j] = fh[j], fh[i]
	fh[i].index = i
	fh[j].index = j
}

func (fh *FoodHeap) Push(x any) {
	n := len(*fh)
	item := x.(*FoodItem)
	item.index = n
	*fh = append(*fh, item)
}

func (fh *FoodHeap) Pop() any {
	old := *fh
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*fh = old[0 : n-1]
	return item
}

func (fh *FoodHeap) Peek() any {
	return (*fh)[0]
}

// update modifies the priority and value of an Item in the queue.
func (fh *FoodHeap) update(item *FoodItem, rating int) {
	item.rating = rating
	heap.Fix(fh, item.index)
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	fr := FoodRatings{
		Cuisines: make(map[string]*FoodHeap),
		Foods:    make(map[string]*FoodItem),
	}
	for i := 0; i < len(foods); i++ {
		item := &FoodItem{
			name:    foods[i],
			cuisine: cuisines[i],
			rating:  ratings[i],
		}
		fr.Foods[item.name] = item
		if _, ok := fr.Cuisines[item.cuisine]; !ok {
			fh := make(FoodHeap, 0)
			heap.Init(&fh)
			fr.Cuisines[item.cuisine] = &fh
		}
		heap.Push(fr.Cuisines[item.cuisine], item)
	}
	return fr
}

func (fr *FoodRatings) ChangeRating(food string, newRating int) {
	item := fr.Foods[food]
	fh := fr.Cuisines[item.cuisine]
	fh.update(item, newRating)
}

func (fr *FoodRatings) HighestRated(cuisine string) string {
	fh := fr.Cuisines[cuisine]
	return fh.Peek().(*FoodItem).name
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */
// @lc code=end
