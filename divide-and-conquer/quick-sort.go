package main

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	} else {
		pivot := nums[0]
		var greaterSub []int
		var lowerSub []int

		for i := range nums[1:] {
			if pivot > nums[i] {
				lowerSub = append(lowerSub, nums[i])
			} else {
				greaterSub = append(greaterSub, nums[i])
			}
		}
		lowerSub = append(lowerSub, []int{pivot}...)
		lowerSub = append(lowerSub, greaterSub...)
		return quickSort(lowerSub)
	}
}
