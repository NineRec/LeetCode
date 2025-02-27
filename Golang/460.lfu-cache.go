package golang

/*
 * @lc app=leetcode id=460 lang=golang
 *
 * [460] LFU Cache
 */

// @lc code=start
// LFUCache with map for o(1) get. list for update counter
import (
	"container/list"
)

type LFUCache struct {
	capacity, minFreq int
	cache             map[int]*list.Element
	freq              map[int]*list.List
}

type Node struct {
	key, value, frequency int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		minFreq:  0,
		cache:    make(map[int]*list.Element, capacity),
		freq:     make(map[int]*list.List, capacity),
	}
}

// Get func, need to cover the frequency change
func (this *LFUCache) Get(key int) int {
	ele, ok := this.cache[key]
	if !ok {
		return -1
	}

	node := ele.Value.(*Node)
	preList := this.freq[node.frequency]
	preList.Remove(ele)

	if preList.Len() == 0 {
		delete(this.freq, node.frequency)

		if this.minFreq == node.frequency {
			this.minFreq++
		}
	}

	node.frequency++
	newList, ok := this.freq[node.frequency]
	if !ok {
		newList = list.Init()
		this.freq[node.frequency] = newList
	}
	newEle := newList.PushBack(node)
	this.cache[key] = newEle

	return node.value
}

// Put func, need to handle update, cap, new
func (this *LFUCache) Put(key int, value int) {
	ele, ok := this.cache[key]
	if ok {
		node := ele.Value.(*Node)
		node.value = value

		this.Get(key)
		return
	}

	if this.capacity == len(this.cache) { // remove 1 ele minFreq
		minList := this.freq[this.minFreq]

		ele := minList.Front()
		minNode := ele.Value.(*Node)
		delete(this.cache, minNode.key)
		minList.Remove(ele)
		if minList.Len() == 0 {
			delete(this.freq, minNode.frequency)
		}
	}

	this.minFreq = 1
	newNode := &Node{
		key:       key,
		value:     value,
		frequency: 1,
	}

	newList, ok := this.freq[newNode.frequency]
	if !ok {
		newList = list.Init()
		this.freq[newNode.frequency] = newList
	}
	newEle := newList.PushBack(newNode)
	this.cache[key] = newEle
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
