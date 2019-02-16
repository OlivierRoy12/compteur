package main

import (
	"fmt"
	"time"
)

func main() {
	for count := 1; count <= 10; count++ {
		if count > 1 {
			fmt.Printf(",")
		}
		fmt.Printf("%d", count)
		time.Sleep(time.Second / 100)
	}
	fmt.Printf("\nGO")
	text("Le gros")
	text("Le petit")

	N1 := 3
	N2 := 2

	addition(N1, N2)
}
