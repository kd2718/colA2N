package convert

/*
* Contains functions to compute column values
 */

import (
	"fmt"
	"math"
	"regexp"
)

var re = regexp.MustCompile("(?i)abcdefghijklmnopqrstuvwxyz")

func reverse_string(runes []rune) []rune {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return runes
}

func ColA2N(alpha []rune) int64 {
	alpha = reverse_string(alpha)
	var total int64 = 0
	var alpha_lookup int64

	for idx, letter := range alpha {
		// get letter value lookup
		alpha_lookup = int64(re.FindStringSubmatchIndex(string(letter))[0]) + 1
		fmt.Println(letter, alpha_lookup)
		fmt.Println(math.Pow(27, float64(idx)))

		total = total + int64(math.Pow(27, float64(idx))) + alpha_lookup

	}

	return total

}
