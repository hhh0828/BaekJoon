package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			if j > i {
				fmt.Print("*")

			} else {
				fmt.Print(" ")
			}
		}
		for j := 5; j >= 0; j-- {
			if j > i {
				fmt.Print("*")

			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()

	}

}
