package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	n := searchInsert(nums, 5)
	fmt.Println(n)

}

func searchInsert(nums []int, target int) int {
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
	return left
}
