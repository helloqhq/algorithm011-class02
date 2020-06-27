package Week_01

func getHint(secret string, guess string) string {
	bucket := make(map[byte]int, 10)
	allMatch, notMatch := 0, 0
	secretBytes := []byte(secret)
	guessBytes := []byte(guess)
	for k, v := range secretBytes {
		if v == guessBytes[k] {
			allMatch++
			continue
		}
		bucket[v]++
		bucket[guessBytes[k]]--
	}
	for _, v := range bucket {
		if v > 0 {
			notMatch += v
		}
	}
	return fmt.Sprintf("%dA%dB", allMatch, len(secretBytes)-notMatch-allMatch)
}
