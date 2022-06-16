type RecentCounter struct {
	i      int
	events []int
	count  int
}

func Constructor() RecentCounter {
	return RecentCounter{
		i:      0,
		events: make([]int, 0),
		count:  0,
	}
}

func (r *RecentCounter) Ping(t int) int {
	r.count++
	r.events = append(r.events, t)
	
	for r.events[r.i] < t - 3000 {
		r.count--
		r.i++
	}
	
	return r.count
}
