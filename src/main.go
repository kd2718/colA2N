package main

import (
	"fmt"
	"strings"

	"github.com/kd2718/colA2N/src/convert"
)

func main() {
	// fmt.Println("Hello from Atom!!!")

	alpha_in := []rune(strings.ToUpper("al"))
	fmt.Println("start test")
	out := convert.ColA2N(alpha_in)

	fmt.Println("ColA2N is", out)

	out2 := convert.ColN2A(out)
	fmt.Println("ColN2A is", out2)
}
