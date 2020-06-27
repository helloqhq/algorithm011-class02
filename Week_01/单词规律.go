package Week_01

import "strings"

func wordPattern(pattern string, str string) bool {
	patternArr, strArr := []byte(pattern), strings.Split(str, " ")
	if len(patternArr) != len(strArr) {
		return false
	}

	patternUsed := make(map[byte]bool)
	strMapPattern := make(map[string]byte)
	newStr := ""
	for k, v := range strArr {
		if _, ok := strMapPattern[v]; !ok {
			if _, ok := patternUsed[patternArr[k]]; ok {
				return false
			}

			strMapPattern[v] = patternArr[k]
			patternUsed[patternArr[k]] = true
		}

		newStr = newStr + string(strMapPattern[v])
	}
	return newStr == pattern
}
