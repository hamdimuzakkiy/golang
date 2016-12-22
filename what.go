package main

func IdontEvenKnowWhatIsThisFunctionDo(a int, b int) int {
	for {
		if a < b {
			return a
		}

		temp := a / b
		a -= (temp * b)
	}
}
