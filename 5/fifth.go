package fifth

func majorityElement(nums []int) int {
	record := map[int]int{}
	majority := int(len(nums) / 2)
	for _, num := range nums {
		if _, ok := record[num]; !ok {
			record[num] = 1
		} else {
			record[num]++
		}
	}

	for key, val := range record {
		if val > majority {
			return key
		}
	}

	return 0
}
