package examples

import "fmt"

func Variables() {
	var a = "initialize" // var declares one or more variables
	var b, c int = 1, 3  // You can declare multiple variables
	var d = false        // Go will infer the type of initialization
	var e int            // Variable declared without a corresponding initialization are non-zero
	f := "apple"         // The := syntax is a shorthand for declaring and initializing a variable

	fmt.Println(a, b, c, d, e, f)
}
