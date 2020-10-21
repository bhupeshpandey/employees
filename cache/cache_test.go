package cache

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	//capacity := 3
	//cache := NewLRU(capacity)
	//
	//// Assertions to check if the hashTable was initialized with the capacity informed
	//assert.NotNil(t,cache)
	//assert.Equal(t, capacity, cache.capacity)
	//
	////Put items in the map/hashTable
	//i :=0 ;
	//fmt.Printf("Starting caching ..\n")
	//for i <= capacity*2 {
	//	key := i
	//	value := i*2
	//	fmt.Printf("Key: %d, Value: %d \n", key, value)
	//	cache.Put(key,value)
	//	i++
	//}
	//// Check if kept just the last entries
	//for i:= capacity*2; i > capacity; i-- {
	//	last := cache.Get(i)
	//	fmt.Printf("Result caching ..\n")
	//	fmt.Printf("Key: %d, Value: %d \n", i, last)
	//	assert.NotNil(t,last)
	//	assert.Equal(t, i*2, last)
	//}
	//
	//// Check if the first one was evicted as should be
	//first := cache.Get(0)
	//assert.NotNil(t,first)
	//assert.Equal(t, -1, first)

}