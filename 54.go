func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	m := len(matrix[0])
	res := make([]int, n*m, n*m)
	flRight := true
	flLeft := false
	flUp := false
	flDown := false
	rightBorder := m - 1
	leftBorder := 0
	downBorder := n - 1
	upBorder := 1
	i := 0
	j := 0
	for k := 0; k < n*m; k++ {
		if flRight {
			if j < rightBorder {
				res[k] = matrix[i][j]
				j++
			} else if j == rightBorder {
				res[k] = matrix[i][j]
				flRight = false
				flDown = true
				rightBorder--
				i++
			}
		} else if flDown {
			if i < downBorder {
				res[k] = matrix[i][j]
				i++
			} else if i == downBorder {
				res[k] = matrix[i][j]
				flDown = false
				flLeft = true
				downBorder--
				j--
			}
		} else if flLeft {
			if j > leftBorder {
				res[k] = matrix[i][j]
				j--
			} else if j == leftBorder {
				res[k] = matrix[i][j]
				flLeft = false
				flUp = true
				leftBorder++
				i--
			}
		} else if flUp {
			if i > upBorder {
				res[k] = matrix[i][j]
				i--
			} else if i == upBorder {
				res[k] = matrix[i][j]
				flUp = false
				flRight = true
				upBorder++
				j++
			}
		}
	}
	return res
}
