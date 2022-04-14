package main

func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	for i < j {
		if numbers[i]+numbers[j] == target {
			break
		}

		if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}

	res := make([]int, 2)
	res[0] = i
	res[1] = j
	return res
}

func main() {

}
