package main

func sum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return nums[0] + sum(nums[1:])
}
