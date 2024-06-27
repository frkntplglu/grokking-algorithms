package main

func sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + sum(nums[1:])
}

func count(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	return 1 + count(nums[1:])

}

func findMax(nums []int, n int) int {
	if n == 1 {
		return nums[0]
	}
	max := findMax(nums, n-1)

	if nums[n-1] > max {
		return nums[n-1]
	}

	return max
}
