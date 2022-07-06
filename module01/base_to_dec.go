package module01

import (
	"fmt"
	"math"
	"strings"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	chars := "0123456789ABCDEF"
	res := 0
	for _, i := range value {
		fmt.Print(strings.IndexRune(chars, i))
		res += int(math.Pow(float64(strings.IndexRune(chars, i)), float64(base)))
	}
	return res
}
