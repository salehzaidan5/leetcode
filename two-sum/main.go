package main

func twoSumQuadratic(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	panic("unreachable")
}

func twoSumLinear(nums []int, target int) []int {
	// Hash map to store the indices of each number in nums.
	indices := make(map[int]int)
	for i, num := range nums {
		indices[num] = i
	}

	// Find the other pair that sum to target in the hash map.
	for i, a := range nums {
		b := target - a
		j, found := indices[b]
		if !found || i == j {
			continue
		}

		return []int{i, j}
	}

	panic("unreachable")
}
