package utils

import (
	"math/rand/v2"
)

func RandomChars(length int) string {
	// A-Z = 65-90 	upper case chars
	// a-z = 97-122 lower case chars
	// ' ' = 95		space
	phrase := ""
	for range length {
		// charSpec Specifies which type of char to use,
		// either lower, upper, or space.
		charSpec := RandomNum(0, 53)
		if charSpec <= 25 {
			randnum := RandomNum(65, 90)
			phrase += string(rune(byte(randnum)))
		} else if charSpec <= 51 {
			randnum := RandomNum(97, 122)
			phrase += string(rune(byte(randnum)))
		} else {
			phrase += " "
		}
	}
	return phrase
}

func RandomNum(start, end int) int {
	end -= start
	randnum := rand.IntN(end)
	return randnum + start
}
