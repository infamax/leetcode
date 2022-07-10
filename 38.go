func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	prevRes := "1"
	var res strings.Builder

	for i := 2; i <= n; i++ {
		res.Reset()
		k := 0
		for k < len(prevRes) {
			j := k
			for j < len(prevRes) && prevRes[j] == prevRes[k] {
				j++
			}


			res.WriteString(strconv.Itoa(j - k))
			res.WriteByte(prevRes[k])
			k = j
		}
		prevRes = res.String()
	}
	return res.String()
}
