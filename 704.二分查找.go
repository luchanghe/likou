package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	index := search(nums, 9)
	fmt.Println(index)
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		min := left + (right-left)/2
		if nums[min] == target {
			return min
		}
		if target < nums[min] {
			right = min - 1
		} else {
			left = min + 1
		}
	}
	return -1
}
