package cache

import (
	"container/list"
)

type LRUCache struct {
	capacity   int
	linkedList *list.List
	hashTable  map[string]*list.Element
}

// The new will initialize the map and linkedList
func NewLRU(capacity int) *LRUCache {
	return &LRUCache{
		capacity:   capacity,
		linkedList: list.New(),
		hashTable:  make(map[string]*list.Element),
	}
}

// Return the value of the key or -1 when it is not found
func (cache *LRUCache) Get(key string) interface{} {
	// Check if exist
	element, has := cache.hashTable[key]
	// if not exist return -1
	if !has {
		return ""
	}
	// if exist move the element before to the first of
	cache.linkedList.MoveBefore(element, cache.linkedList.Front())

	// each element is the linked  list pointer where []int[0] is the index and the []int[1] is the value
	return element.Value
}

func (cache *LRUCache) Put(key string, value interface{}) {
	element, has := cache.hashTable[key]
	if has {
		// if has element with the key then add the new value
		element.Value = value
		// move element to the front of the list
		cache.linkedList.MoveBefore(element, cache.linkedList.Front())
		return
	}
	if len(cache.hashTable) >= cache.capacity {
		Evict(cache)
	}
	createElementOnFrontOfList(cache, key, value)
	return

}

// Create a new element in front of the list
func createElementOnFrontOfList(cache *LRUCache, key string, value interface{}) {
	front := cache.linkedList.PushFront([]interface{}{key, value})
	cache.hashTable[key] = front
}

// Remove the last item of the map
func Evict(cache *LRUCache) {
	back := cache.linkedList.Back()
	cache.linkedList.Remove(back)
	delete(cache.hashTable, back.Value.([]string)[0])
}

