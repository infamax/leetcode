func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			f := i
			s := 0
			for f < len(haystack) && s < len(needle) && haystack[f] == needle[s] {
				f++
				s++
			}

			if s == len(needle) {
				return i
			}
		}
	}

	return -1
}
