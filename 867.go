func transpose(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])
	res := make([][]int, m, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = matrix[j][i]
		}
	}
	return res
}
