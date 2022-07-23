package main

import (
	"fmt"
)

func main() {

	var j int

	for j = 1; j <= 100; j = j + 1 {

		if j%2 == 0 {
			fmt.Println("even number: ", j)
		}

	}
}
