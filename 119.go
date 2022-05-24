
func getPascalTriangle(num int) [][]int {
	res := make([][]int, num+1, num+1)
	res[0] = make([]int, 1, 1)
	res[0][0] = 1
	count := 2
	for i := 1; i < len(res); i++ {
		res[i] = make([]int, count, count)
		res[i][0] = 1
		res[i][count-1] = 1
		if count == 2 {
			count++
			continue
		}
		for j := 1; j < count-1; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}
		count++
	}
	return res
}

func getRow(rowIndex int) []int {
	res := getPascalTriangle(rowIndex)
	return res[rowIndex]
}
