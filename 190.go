func positionsInZeroes(num uint32) [32]bool {
	var res[32]bool
	for i := 0; i < 32; i++ {
		if num & (1 << i) > 0 {
			res[i] = true
		}
	}
	return res
}


func reverseBits(num uint32) uint32 {
	positionsZeroes := positionsInZeroes(num)
	var res uint32 = 0
	for i := 0; i < 32; i++ {
		if positionsZeroes[i] {
			res |= 1 << (31 - i)
		}
	}	
	return res
}
