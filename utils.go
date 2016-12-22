package main

import (
	"errors"
	"fmt"
)

func ReverseString(s string) string {
	if len(s) == 0 {
		return ""
	}
	return fmt.Sprintf("%c%s", s[len(s)-1], ReverseString(s[:len(s)-1]))
}

func Power(a, b int) int {
	temp := 1
	for b > 0 {
		temp *= a
		b -= 1
	}
	return temp
}

func AbsoultDivisor(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Divisor by 0")
	}
	if a/b < 0 {
		return -a / b, nil
	}
	return a / b, nil
}
