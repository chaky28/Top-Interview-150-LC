package second

func removeElement(nums []int, val int) int {
	repetition := len(nums)
	iter := 0
	for iter < repetition {
		if nums[iter] == val {
			for j := iter; j < repetition-1; j++ {
				nums[j] = nums[j+1]
			}
			repetition--
		} else {
			iter++
		}
	}
	return repetition - 1
}
