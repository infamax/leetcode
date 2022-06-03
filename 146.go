const inf = math.MaxInt - 1

type LRUCache struct {
	data        map[int]int
	used        map[int]int
	length      int
	capacity    int
	numLastItem int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		data:        make(map[int]int),
		used:        make(map[int]int),
		length:      0,
		capacity:    capacity,
		numLastItem: -1,
	}
}

func (l *LRUCache) Get(key int) int {
	if val, ok := l.data[key]; ok {
		l.numLastItem++
		l.used[key] = l.numLastItem
		return val
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	l.numLastItem++
	if _, ok := l.data[key]; ok {
		l.data[key] = value
		l.used[key] = l.numLastItem
		return
	}

	if l.length < l.capacity {
		l.length++
		l.data[key] = value
		l.used[key] = l.numLastItem
		return
	}
	deleteKey := findMinimalUsedKey(l.used)
	delete(l.data, deleteKey)
	delete(l.used, deleteKey)
	l.data[key] = value
	l.used[key] = l.numLastItem
}

func findMinimalUsedKey(used map[int]int) int {
	ans := -1
	min := inf
	for key, value := range used {
		if value < min {
			min = value
			ans = key
		}
	}
	return ans
}
