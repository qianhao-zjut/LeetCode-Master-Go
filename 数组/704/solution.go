package main

import "fmt"

//704.二分查找
func search(nums []int, target int) int {
	begin := 0
	end := len(nums)
	for begin < end {
		mid := (begin + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			begin = mid + 1
		} else {
			end = mid
		}
	}
	return -1
}






func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, -1))
}
