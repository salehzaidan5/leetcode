package main

func searchInsert(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		m := i + (j-i)/2
		if nums[m] > target {
			j = m - 1
		} else if nums[m] < target {
			i = m + 1
		} else {
			return m
		}
	}

	return i
}
