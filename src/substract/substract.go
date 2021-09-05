package substract

func Substract(x, y int) int {
	return x - y
}

func Absolute(value int) int {
	if value < 0 {
		return 0 - value
	}
	return value
}
