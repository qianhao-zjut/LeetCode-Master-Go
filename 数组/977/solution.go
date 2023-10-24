package main



func sortedSquares(nums []int) []int {

	maxIndex := len(nums) - 1
	minIndex := 0

	res := make([]int, 0,len(nums))
	res_index := len(nums) - 1

	for minIndex <= maxIndex {

		if square(nums[maxIndex]) > square(nums[minIndex]) {
			res[res_index] = square(nums[maxIndex])
			maxIndex--
		} else {
			res[res_index] = square(nums[minIndex])
			minIndex++
		}

		res_index--
	}
	
	return res
}

func square(x int) int {
    return x * x
}