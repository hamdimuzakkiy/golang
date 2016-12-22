package main

func AbsDivisorAndDivAgain(a, b, c int) {
	div, err := Divisor(10, 2)
	if err != nil {
		return 0
	}
	return c / div
}
