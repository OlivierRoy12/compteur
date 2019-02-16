package main

import (
	"fmt"
	"time"
)

func main() {
	for count := 1; count <= 100; count++ {
		if count > 1 {
			fmt.Printf(",")
		}
		fmt.Printf("%d", count)
		time.Sleep(time.Second / 100)
	}
	fmt.Printf("\nGO")
}
