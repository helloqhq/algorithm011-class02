package Week_06

func leastInterval(tasks []byte, n int) int {
	if len(tasks) == 0 {
		return 0
	}

	max, ele := Count(tasks)
	point := (max-1)*(n+1) + len(ele)
	if point < len(tasks) {
		return len(tasks)
	} else {
		return point
	}
}

func Count(b []byte) (int, []byte) {
	/*定义一个字典统计每个字符的出现次数*/
	m := make(map[byte]int)
	for _, v := range b {
		/*若字典中没有v这个键，则m[v]的值默认为0*/
		m[v]++
	}
	/*把所有出现次数存储到一个切片内*/
	s := make([]int, 0)
	for _, v := range m {
		s = append(s, v)
	}
	/*找到所有出现次数中最大的那一个*/
	maxAccount := s[0]
	for _, v := range s[1:] {
		if v > maxAccount {
			maxAccount = v
		}
	}
	/*遍历字典中的值，如果和maxAccount相等，就把它存进maxElements*/
	maxElements := make([]byte, 0)
	for k, v := range m {
		if v == maxAccount {
			maxElements = append(maxElements, k)
		}
	}
	return maxAccount, maxElements
}
