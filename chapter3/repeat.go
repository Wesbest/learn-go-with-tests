package iteration

import "strconv"

// const repeatCount = 5

func Repeat(character string, amount string) string {
	var repeated string
	j, _ := strconv.Atoi(amount)
	for i := 0; i < j; i++ {
		repeated += character
	}
	return repeated
}
