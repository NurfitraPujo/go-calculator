package absolute

func Absolute(value int) int {
	if value < 0 {
		return 0 - value
	}
	return value
}
