package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Printf("%v\n", result)
}

func twoSum(nums []int, target int) []int {
	hMap := make(map[int]int)
	for k, v := range nums {
		y := target - v
		if _, ok := hMap[y]; !ok {
			hMap[v] = k
		} else {
			return []int{hMap[y], k}
		}
	}
	return nil
}
