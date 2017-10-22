package dictionary

import (
	"strings"
)

// MemoryCache ...
type MemoryCache struct {
	dict map[string]WordData
}

// NewMemoryCache ...
func NewMemoryCache() MemoryCache {
	cache := MemoryCache{}
	cache.dict = make(map[string]WordData)
	return cache
}

// Get ...
func (cache MemoryCache) Get(word string) *WordData {
	if val, ok := cache.dict[strings.Trim(word, " ")]; ok {
		return &val
	}
	return nil
}

// Add ...
func (cache MemoryCache) Add(data WordData) WordData {
	cache.dict[strings.Trim(data.Word, " ")] = data
	return data
}

// Update ...
func (cache MemoryCache) Update(word, explanation string) *WordData {
	if val, ok := cache.dict[word]; ok {
		val.Explanation = &explanation
		return &val
	}
	return nil
}

// Delete ...
func (cache MemoryCache) Delete(word string) bool {
	if _, ok := cache.dict[word]; ok {
		delete(cache.dict, word)
		return true
	}
	return false
}
