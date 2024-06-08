package arrslice

func Sum(n []int) int {
	var sum int
	for _, v := range n {
		sum += v
	}
	return sum
}
func SumAll(slices ...[]int) []int {
	// get the length of the parameter slices
	lenOfNum := len(slices)
	// make a new []int slice with starting capacity of lenOfNum
	sums := make([]int, lenOfNum)
	// sum the currect slice in the loop and put it to the sums[i]
	for i, slice := range slices {
		sums[i] = Sum(slice)
	}
	return sums
}
func SumAllBench1(slices ...[]int) []int {
	var result []int
	for _, slice := range slices {
		total := Sum(slice)
		result = append(result, total)
	}
	return result
}

func SumAllOptimized(slices ...[]int) []int {
	lenOfNum := len(slices)
	sums := make([]int, lenOfNum)

	for i, slice := range slices {
		sum := 0
		for _, num := range slice {
			sum += num
		}
		sums[i] = sum
	}

	return sums
}

func SumAllTails(slices ...[]int) []int {

	lenOfNum := len(slices)
	sums := make([]int, lenOfNum)

	for i, slice := range slices {
		sum := 0
		for idx, num := range slice {
			if idx == 0 {
				continue
			}
			sum += num

		}
		sums[i] = sum
	}
	return sums

}
