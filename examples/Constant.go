package examples

import (
	"fmt"
	"math"
)

const s string = "constant" // const declared a constant value
func Contant() {
	fmt.Println(s)
	const n = 50000000 // A constant statement can appear anywhere a var statement can
	const d = 3e20 / n // A constant expression perform arithmetic with arbitrary
	fmt.Println(int64(d))
	fmt.Println(math.Sin(d))

}
