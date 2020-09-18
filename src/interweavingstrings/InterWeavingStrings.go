package interweavingstrings

func InterweavingStrings(one, two, three string) bool {
	if len(three) != len(one)+len(two) {
		return false
	}
	a, b := 0, 0
	for i := 0; i < len(three); i++ {
		//tc := three[i]
		// fmt.Printf("third String character:%c",byte(tc))
		// fmt.Println()
		if a < len(one) && b < len(two) && one[a] == three[i] && two[b] == three[i] {
			if !InterweavingStringsHelper(a+1, b, i+1, one, two, three) {
				if !InterweavingStringsHelper(b+1, a, i+1, two, one, three) {
					return false
				}
				return true
			}
			return true
		}
		if a < len(one) && one[a] == three[i] {
			a++
		} else if b < len(two) && two[b] == three[i] {
			b++
		}
		// fmt.Printf("first string count:%d",a)
		// fmt.Println()
		// fmt.Printf("Second string count:%d",b)
		// fmt.Println()
	}

	return a == len(one) && b == len(two)
}

func InterweavingStringsHelper(a, b, i int, one, two, three string) bool {

	for ; i < len(three); i++ {
		// tc := three[i]
		// fmt.Printf("third String character:%c",byte(tc))
		// fmt.Println()
		if a < len(one) && b < len(two) && one[a] == three[i] && two[b] == three[i] {
			if !InterweavingStringsHelper(a+1, b, i+1, one, two, three) {
				if !InterweavingStringsHelper(b+1, a, i+1, two, one, three) {
					return false
				}
				return true
			}
			return true
		}
		if a < len(one) && one[a] == three[i] {
			a++
		} else if b < len(two) && two[b] == three[i] {
			b++
		}
		// fmt.Printf("first string count:%d",a)
		// fmt.Println()
		// fmt.Printf("Second string count:%d",b)
		// fmt.Println()
	}
	return a == len(one) && b == len(two)
}
