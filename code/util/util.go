package util

func Abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func Min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
