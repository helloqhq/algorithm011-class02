package Week_01

func plusOne(digits []int) []int {
	l := len(digits) - 1
	for i := l; i >= 0; i-- {
		digits[i] += 1
		if digits[i] <= 9 {
			return digits
		} else {
			digits[i] = 0
		}
	}
	digits[0] = 1
	digits=append(digits,0)
	return digits
}
