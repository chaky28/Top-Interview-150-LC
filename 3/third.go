package third

func removeDuplicates(nums []int) int {
	maxLen := len(nums)
	iter := 0
	for iter < maxLen {
		if nums[iter] == nums[iter+1] {
			for j := iter + 1; j < maxLen; j++ {
				nums[j] = nums[j+1]
			}
			maxLen--
		} else {
			iter++
		}
	}

	return maxLen
}
