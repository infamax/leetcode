func setZeroes(matrix [][]int) {
	if len(matrix) == 0  {
		return
	}
	
	rowsSetZeroes := make([]bool, len(matrix), len(matrix))
	columnsSetZeroes := make([]bool, len(matrix[0]), len(matrix[0]))
	
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rowsSetZeroes[i] = true
				columnsSetZeroes[j] = true
			}
		} 
	}
	
	for i := 0; i < len(rowsSetZeroes); i++ {
		if rowsSetZeroes[i] {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	
	for j := 0; j < len(columnsSetZeroes); j++ {
		if columnsSetZeroes[j] {
			for i := 0; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}
}
