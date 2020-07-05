package Week_02

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v]++
	}
	k := 0
	for _, v := range nums2 {
		if m[v] > 0 {
			m[v]--
			nums1[k] = v
			k++
		}
	}
	return nums1[0:k]
}