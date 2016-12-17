package main

import (
	"fmt"
)

func ReverseString(s string) string {
	if len(s) == 0 {
		return ""
	}
	return fmt.Sprintf("%c%s", s[len(s)-1], ReverseString(s[:len(s)-1]))
}
