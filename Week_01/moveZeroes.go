package Week_01

func moveZeroes(nums []int) {
	j := 0
	for i, v := range nums {
		if v != 0 {
			if i != j {
				nums[j] = v
				nums[i] = 0
			}
			j++
		}
	}
}


