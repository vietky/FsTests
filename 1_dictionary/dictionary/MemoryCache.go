package dictionary

// MemoryCache ...
type MemoryCache struct {
	dict map[string]WordData
}

// Init ...
func (cache MemoryCache) Init() {
	cache.dict = make(map[string]WordData)
}

// Get ...
func (cache MemoryCache) Get(word string) *WordData {
	if val, ok := cache.dict[word]; ok {
		return &val
	}
	return nil
}

// Add ...
func (cache MemoryCache) Add(data WordData) WordData {
	cache.dict[data.Word] = data
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
