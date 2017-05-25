package main

import (
	// "errors"
	"fmt"
)

func ReverseString(s string) string {
	if len(s) == 0 {
		return ""
	}
	return fmt.Sprintf("%c%s", s[len(s)-1], ReverseString(s[:len(s)-1]))
}

func Power(a, b int) int {
	if b == 0 {
		return 1
	}

	temp := Power(a, b/2)

	if b%2 == 0 {
		return temp * temp
	} else {
		return temp * temp * a
	}
}

func Plus(a, b int) int {
	return a + b
}

func AbsoultDivisor(a, b int) (int, error) {
	if b == 0 {
		return 0, nil
	}
	if a/b < 0 {
		return -a / b, nil
	}
	return a / b, nil
}

func Sub(a, b int) int {
	return a - b
}
