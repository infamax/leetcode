func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	if len(arr) < 3 {
		return -1
	}

	count := 0

	for i := 0; i < len(arr) - 2; i++ {
		for j := i + 1; j < len(arr) - 1; j++ {
			for k := j + 1; k < len(arr); k++ {
				if abs(arr[i] - arr[j]) <= a &&
					abs(arr[j] - arr[k]) <= b && 
					abs(arr[i] - arr[k]) <= c {
					count++
				}
			}
		}
	}
	return count
}
