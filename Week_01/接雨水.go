package Week_01

import "math"

func trap(height []int) int {
	sum := 0
	heightLen := len(height)
	maxL, maxR := make([]float64, heightLen), make([]float64, heightLen)
	for i := heightLen - 2; i >= 0; i-- {
		maxR[i] = math.Max(maxR[i+1], float64(height[i+1]))
	}
	for i, v := range height {
		if i == 0 {
			maxL[i] = 0
		} else {
			maxL[i] = math.Max(maxL[i-1], float64(height[i-1]))
		}

		min := math.Min(maxL[i], maxR[i])
		if min > float64(v) {
			sum = sum + int(min) - v
		}
	}
	return sum
}
