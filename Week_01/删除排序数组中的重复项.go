package Week_01

func removeDuplicates(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	i := 0
	for j := 1; j < numsLen; j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
