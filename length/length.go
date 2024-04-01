package length

func StringLength(s string) int {
	return len([]rune(s))
}
