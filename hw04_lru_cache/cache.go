package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	mu       *sync.Mutex
	queue    List
	items    map[Key]*ListItem
}

type CacheItem struct {
	Key   Key
	Value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		mu:       &sync.Mutex{},
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (lru *lruCache) Set(key Key, value interface{}) bool {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	if el, ok := lru.items[key]; ok {
		lru.queue.MoveToFront(el)
		el.Value = value
		return true
	}
	if lru.queue.Len() == lru.capacity {
		purge(lru)
	}
	node := lru.queue.PushFront(value)
	lru.items[key] = node
	return false
}

func (lru *lruCache) Get(key Key) (interface{}, bool) {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	if el, ok := lru.items[key]; ok {
		lru.queue.MoveToFront(el)
		return el.Value, true
	}
	return nil, false
}

func (lru *lruCache) Clear() {
	lru.mu.Lock()
	lru.items = make(map[Key]*ListItem, lru.capacity)
	lru.mu.Unlock()
}

func purge(lru *lruCache) {
	if el := lru.queue.Back(); el != nil {
		lru.queue.Remove(el)
		delete(lru.items, el.Key)
	}
}
