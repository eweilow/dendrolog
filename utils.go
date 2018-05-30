package dendrolog

// Repeat a string
func repeat(times int, substr string) string {
	var s string
	for i := 0; i < times; i++ {
		s += substr
	}
	return s
}

// Return max of a and b
func intMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// Returns the string in strs with maximum length
func maxStringLength(strs []string) int {
	max := 0
	for i := range strs {
		strLen := len(strs[i])
		if strLen > max {
			max = strLen
		}
	}
	return max
}
