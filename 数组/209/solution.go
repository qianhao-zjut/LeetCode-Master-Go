package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {

	var beginIndex, endIndex, res, sum int
	sum = nums[0]
	for ;; {
		if sum >= target {
			if res == 0 {
				res = endIndex - beginIndex + 1
			} else {
				res = min(res, endIndex-beginIndex+1)
			}
			sum = sum - nums[beginIndex]
			beginIndex++
			if beginIndex >= len(nums){
				break
			}
		} else {
			endIndex++
			if endIndex >= len(nums){
				break
			}
			sum = nums[endIndex] + sum
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	array := []int{2, 3, 1, 2, 4, 3}
	res := minSubArrayLen(7, array)
	fmt.Println(res)
}
