package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	fmt.Println(5)
	for i := 0; i < 10; i++ {

		for j := 0; i < 10; i++ {
			if j < i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()

	}
}
