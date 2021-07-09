package main

import "fmt"

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	n := sortedSquares(nums)
	fmt.Println(n)
}
func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	l, r := 0, n-1
	for pos := n - 1; pos >= 0; pos-- {
		left, right := nums[l]*nums[l], nums[r]*nums[r]
		if left > right {
			ans[pos] = left
			l++
		} else {
			ans[pos] = right
			r--
		}
	}
	return ans

}
