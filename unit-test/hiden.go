package main

func AbsDivisorAndDivAgain(a, b, c int) int {
	div, err := AbsoultDivisor(10, 2)
	if err != nil {
		return 0
	}
	return c / div
}
