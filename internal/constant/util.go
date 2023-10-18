package constant

import "math/rand"

func RandomSixDigitNumber() int {
	min := 100000
	max := 999999
	return rand.Intn(max-min+1) + min
}
