package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0
	right := 0
	for {
		for left < len(nums) && nums[left] != val {
			left++
		}

		for right < len(nums) && (right <= left || nums[right] == val) {
			right++
		}
		if right == len(nums) {
			break
		}

		nums[left], nums[right] = nums[right], nums[left]
	}
	return left
}
