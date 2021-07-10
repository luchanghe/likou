package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	n := twoSum(nums, target)
	fmt.Printf("%v", n)
}

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{-1, -1}
}
