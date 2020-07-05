package Week_02

func maxSlidingWindow(nums []int, k int) []int {
	numsLen := len(nums)
	if numsLen == 0 {
		return []int{}
	}

	deque := make([]int, 0, k)
	for i := 0; i < k; i++ {//遍历小于滑动窗口数据，压入队列，并保持队列左侧值最大
		for len(deque) > 0 && deque[len(deque)-1] < nums[i] {//清除队列中小于当前值的数据
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, nums[i])
	}

	res := make([]int, 0, numsLen-k+1)
	res = append(res, deque[0])//首个滑动窗口最大值

	for i := k; i < numsLen; i++ {
		if nums[i-k] == deque[0] {//队列里最大值，即队列头，不在当前滑动窗口范围内时，弹出队列头抛弃
			deque = deque[1:]
		}

		for len(deque) > 0 && deque[len(deque)-1] < nums[i] {//清除队列中小于当前值的数据
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, nums[i])
		res = append(res, deque[0])
	}
	return res
}
