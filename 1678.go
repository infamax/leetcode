func interpret(command string) string {
	res := make([]rune, 0)
	count := 0
	for _, v := range command {
		if v == 'G' {
			res = append(res, v)
			continue
		}

		if v == '(' {
			continue
		}

		if v == ')' {
			if count == 0 {
				res = append(res, 'o')
			} else {
				res = append(res, 'a', 'l')
			}
			count = 0
			continue
		}

		count++
	}
	return string(res)
}
