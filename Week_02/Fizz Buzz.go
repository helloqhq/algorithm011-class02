package Week_02

import "strconv"

func fizzBuzz(n int) []string {
	dictKey := []int{3, 5}
	dictVal := []string{"Fizz", "Buzz"}
	var ans []string
	for i := 1; i <= n; i++ {
		subAns := ""
		for k, v := range dictKey {
			if i%v == 0 {
				subAns = subAns + dictVal[k]
			}
		}
		if subAns != "" {
			ans = append(ans, subAns)
		} else {
			ans = append(ans, strconv.Itoa(i))
		}
	}
	return ans
}
