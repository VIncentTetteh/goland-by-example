package examples

import "fmt"

func ForLoop() {
	// for is Gos only looping construct.
	i := 1
	for i <= 2 {
		fmt.Print(i)
		i += 1
	}

	for j := 0; j <= 10; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for i := 0; i < 11; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
