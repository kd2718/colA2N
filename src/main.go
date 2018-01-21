package main

import (
	"fmt"

	"github.com/kd2718/colA2N/src/convert"
)

func main() {
	// fmt.Println("Hello from Atom!!!")

	alpha_in := []rune("c")
	out := convert.ColA2N(alpha_in)

	fmt.Println(out)
}
