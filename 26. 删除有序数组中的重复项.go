package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 2, 3, 3, 3, 3, 3, 5, 5, 5, 5, 6, 6, 6, 6, 7}
	p := removeDuplicates(nums)
	fmt.Printf("%d", p)
}

func removeDuplicates(nums []int) int {
	down := 1
	for top := 1; top < len(nums); top++ {
		if nums[top] != nums[top-1] {
			nums[down] = nums[top]
			down++
		}
	}
	return down
}
