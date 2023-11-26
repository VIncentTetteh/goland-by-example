package examples

import "fmt"

func IFElse() {
	if 7%2 == 0 {
		fmt.Println(" 7 is even")
	} else {
		fmt.Println("7 is old")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 3 scissors")
	}

	if num := 9; num < 0 {
		fmt.Println(num, " is nagative")
	} else if num < 10 {
		fmt.Println(num, " has one digit")
	} else {
		fmt.Println(num, " has multiple of digit")
	}
}
