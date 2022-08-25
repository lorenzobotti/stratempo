package stratempo

func powInt(base, exp int) int {
	if exp == 0 {
		return 1
	}

	res := base
	for i := 1; i < exp; i++ {
		res *= base
	}

	return res
}
