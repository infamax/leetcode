func myAtoi(s string) int {
	s = strings.TrimSpace(s)
    if len(s) == 0 {
        return 0
    }
    
	signed := false
    plus := false
	if s[0] == '-' {
		signed = true
	}
    
    if s[0] == '+' {
        plus = true
    }

	i := 0
    
    if signed && plus {
        return 0
    }

    if signed || plus {
		i = 1
	}

	str := []rune(s)
	for i < len(str) && str[i] >= '0' && str[i] <= '9' {
		i++
	}

	if !signed && i == 0 {
		return 0
	}

	if signed && i == 1 {
		return 0
	}

	str = str[:i]
	i = 0

    if signed || plus {
		i = 1
	}

	for i < len(str) && str[i] == '0' {
		i++
	}

	countRanks := 1
	res := 0
    count := 0
	for j := len(str) - 1; j >= i; j-- {
		res += int(str[j]-'0') * countRanks
		countRanks *= 10
        count++
	}
    
    if count > 10 && !signed {
        return 2147483647
    }
    
    if count > 10 && signed {
        return -2147483648
    }

    if abs(res) > 2147483647 && !signed {
		return 2147483647
	} 
	
    if abs(res) > 2147483648 && signed {
		return -2147483648
	}

	if signed {
		return -res
	}
    
	return res
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
