func countPrimes(n int) int {
	a := make([]bool, n, n)
	for i := 0; i < len(a); i++ {
		a[i] = true
	}
	count := 0
	for i := 2; i < len(a); i++ {
		if a[i] == true {
			count++
			for j := i; j < len(a); j += i {
				a[j] = false
			} 
		}
 	} 	
	return count
}
