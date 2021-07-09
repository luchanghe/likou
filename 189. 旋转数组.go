package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(nums)
}
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	run(nums[:n-k])
	run(nums[n-k:])
	run(nums)
}

func run(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
