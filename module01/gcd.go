package module01

func GCD(a, b int) int {
	for b != 0{
		a, b = b, a%b
	}
	return a
}

// 30 9 -> 3
// 30 / 9
