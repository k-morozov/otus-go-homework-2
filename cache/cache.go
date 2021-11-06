package cache

import "sync"

type Cache interface {
	Set(key string, value interface{}) bool
	Get(key string) (interface{}, bool)
	Clear()
}

type LruCache struct {
	size     int
	elements ForwardList
	mp       map[string]*Node
	mu       sync.Mutex
}

func NewCache(size int) Cache {
	return &LruCache{size: size, elements: NewList(), mp: make(map[string]*Node)}
}

func (lru *LruCache) Set(key string, value interface{}) (ok bool) {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	node, ok := lru.mp[key]
	if ok {
		node.Value = value
		lru.elements.MoveToFront(node)
	} else {
		node = lru.elements.PushFront(value)
		lru.mp[key] = node
		if lru.size < lru.elements.Size() {
			node = lru.elements.Back()
			lru.elements.Remove(node)
		}
	}
	return
}

func (lru *LruCache) Get(key string) (interface{}, bool) {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	_, ok := lru.mp[key]
	if ok {
		return lru.mp[key].Value, true
	}
	return nil, false
}

func (lru *LruCache) Clear() {
	// @TBD
}
