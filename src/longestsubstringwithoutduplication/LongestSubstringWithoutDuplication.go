package longestsubstringwithoutduplication

func LongestSubstringWithoutDuplication(str string) string {
	// Write your code here.
	strMap := make(map[byte]int)

	i, j := 0, 0
	longestSubString := ""

	for j < len(str) {
		c := str[j]
		_, ok := strMap[c]
		if ok {
			tempStr := str[i:j]
			if len(tempStr) > len(longestSubString) {
				longestSubString = tempStr
			}
			for i < j {
				t := str[i]
				delete(strMap, t)
				if t == c {
					i++
					break
				}
				i++
			}
		}
		strMap[c] = j

		j++
	}
	tempStr := str[i:j]
	if len(tempStr) > len(longestSubString) {
		longestSubString = tempStr
	}

	return longestSubString
}
