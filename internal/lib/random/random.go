package random

import (
	"math/rand/v2"
)

func NewRandomString(lenght int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	alias := make([]byte, lenght)
	for n := range alias {
		alias[n] = letters[rand.IntN(len(letters))]
	}

	return string(alias)
}

func NewRandomNum(fnum int, snum int) int {

	n := rand.IntN(snum) + fnum

	return n
}
