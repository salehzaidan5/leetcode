package main

func removeDuplicates(nums []int) int {
	left := 1
	right := 1
	last := nums[0]
	for {
		for right < len(nums) && nums[right] <= last {
			right++
		}
		if right == len(nums) {
			break
		}

		nums[left], nums[right] = nums[right], nums[left]
		last = nums[left]
		left++
	}
	return left
}
