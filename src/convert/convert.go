package convert

/*
* Contains functions to compute column values
 */

import (
	"math"

	"github.com/kd2718/colA2N/src/alphabet"
)

var re = alphabet.AlphaRune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

const numLetters int64 = 26

// var alphamap map[rune]int64 =

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
		alpha_lookup = re.FirstIdx(letter) + 1

		// total is current total added to the power of 26 times the current
		//alpha lookup
		total = total + (int64(math.Pow(float64(numLetters), float64(idx))) * alpha_lookup)
	}

	return total

}

func ColN2A(total int64) string {
	// takes an int and finds the alpha column number
	// num := total
	var rem int64
	var outRn []rune

	for num := total; num > 0; num = num / numLetters {
		rem = num % numLetters
		outRn = append(outRn, re[rem-1])
	}

	outRn = reverse_string(outRn)
	return string(outRn)

}
