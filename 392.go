func isSubsequence(s string, t string) bool {
	if len(t) == 0 && len(s) == 0 {
		return true
	}

	if len(t) == 0 {
		return false
	}
	
	if len(s) == 0 {
		return true
	}

	first, second := 0, 0

	for first < len(s) && second < len(t) {
		if s[first] == t[second] {
			if first == len(s)-1 {
				return true
			}
			first++
			second++
		} else {
			second++
		}
	}
	return false
}
