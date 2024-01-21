package fourth

func removeDuplicates(nums []int) int {
	maxLen := len(nums)
	iter := 0
	for iter < maxLen-2 {
		if nums[iter] == nums[iter+1] && nums[iter+1] == nums[iter+2] {
			for j := iter + 2; j < maxLen-1; j++ {
				nums[j] = nums[j+1]
			}
			maxLen--
		} else {
			if nums[iter] != nums[iter+1] {
				iter++
			} else {
				iter += 2
			}
		}
	}

	return maxLen
}
