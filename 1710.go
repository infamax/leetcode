func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	res := 0

	for i := 0; i < len(boxTypes); i++ {
		if truckSize == 0 {
			break
		}

		if boxTypes[i][0] > truckSize {
			res += boxTypes[i][1] * truckSize
			break
		} else {
			res += boxTypes[i][1] * boxTypes[i][0]
			truckSize -= boxTypes[i][0]
		}
	}

	return res
}
