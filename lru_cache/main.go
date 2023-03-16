package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Data int
	Ptr  *list.Element
}

type LRUCache struct {
	Queue    *list.List
	Items    map[int]*Node
	Capacity int
}

func main() {
	fmt.Println("Test case 1")
	obj := NewLRUCache(4)
	obj.Push(1, 1)
	obj.Push(2, 2)
	obj.Push(6, 7)
	fmt.Println(obj.Fetch(1))
	obj.Push(3, 3)
	fmt.Println(obj.Fetch(2))
	obj.Push(4, 123)
	fmt.Println(obj.Fetch(1))
	fmt.Println(obj.Fetch(3))
	fmt.Println(obj.Fetch(4))
	fmt.Println(obj.Fetch(6))
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{Queue: list.New(), Items: make(map[int]*Node), Capacity: capacity}
}
func (l *LRUCache) Push(key int, value int) {
	item, found := l.Items[key]
	if !found {
		if l.Capacity == len(l.Items) {
			back := l.Queue.Back()
			l.Queue.Remove(back)
			delete(l.Items, back.Value.(int))
		}
		l.Items[key] = &Node{Data: value, Ptr: l.Queue.PushFront(key)}
	} else {
		item.Data = value
		l.Items[key] = item
		l.Queue.MoveToFront(item.Ptr)
	}
}

func (l *LRUCache) Fetch(key int) int {
	item, found := l.Items[key]
	if found {
		l.Queue.MoveToFront(item.Ptr)
		return item.Data
	}
	return -1
}
