package six

func rotate(nums []int, k int) {
	nums2 := make([]int, len(nums))
	copy(nums2, nums)
	for i := 0; i < len(nums); i++ {
		if i+k >= len(nums) {
			pos := i + k - len(nums)
			for pos >= len(nums) {
				pos -= len(nums)
			}
			nums[pos] = nums2[i]
		} else {
			nums[i+k] = nums2[i]
		}
	}
}
