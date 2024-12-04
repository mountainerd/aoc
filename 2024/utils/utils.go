package utils

// AbsDiff implements a rudimentary absolute value function but for integers, not floats.
func AbsDiff[N int | float64](x, y N) N {
	if x < y {
		return y - x
	}

	return x - y
}
