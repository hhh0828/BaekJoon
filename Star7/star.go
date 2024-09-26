package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 2XN-1 번째 줄까지 출력
func main() {
	defer Writer.Flush()
	N := Readint()

	lines := 2*N - 1
	mid := lines / 2

	// 첫 번째 부분: 별이 늘어나는 삼각형
	for i := 1; i <= mid+1; i++ {
		for j := 0; j < lines; j++ {
			if (mid-i < j && j < mid+i) || j == mid {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}

	// 두 번째 부분: 별이 줄어드는 삼각형
	for i := 0; i < mid; i++ {
		for j := 0; j < mid; j++ {
			if j > i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		for j := mid; j >= 0; j-- {
			if j > i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		// 마지막 줄에서는 개행 문자를 출력하지 않음
		if i != mid-1 {
			fmt.Print("\n")
		}
	}
}

var (
	Reader *bufio.Reader
	Writer *bufio.Writer
)

func init() {
	Reader = bufio.NewReaderSize(os.Stdin, 30)
	Writer = bufio.NewWriter(os.Stdout)
}

func Readint() int {
	Rb, _, _ := Reader.ReadLine()
	a, _ := strconv.Atoi(string(Rb))
	return a
}
