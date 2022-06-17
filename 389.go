func makeMapByString(s string) map[rune]int {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	return m
}

func findTheDifference(s string, t string) byte {
	m1 := makeMapByString(s)
	m2 := makeMapByString(t)
	
	var res byte
	
	for key, value := range m2 {
		c, ok := m1[key]
		if !ok {
			res = byte(key)
			break
		}
		
		if c < value {
			res = byte(key)
			break
		}
	}
	return res
}
