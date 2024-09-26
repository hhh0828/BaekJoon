package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	defer Writer.Flush()
	N := Readint()

	lines := 2*N - 1
	mid := lines / 2

	// 첫 번째 부분: 별이 늘어나는 삼각형
	for i := 0; i <= mid; i++ {
		// 왼쪽 공백 출력
		for j := 0; j < mid-i; j++ {
			Writer.WriteString(" ")
		}
		// 별 출력
		for j := 0; j < 2*i+1; j++ {
			Writer.WriteString("*")
		}
		Writer.WriteString("\n")
	}

	for i := mid - 1; i >= 0; i-- {

		for j := 0; j < mid-i; j++ {
			Writer.WriteString(" ")
		}

		for j := 0; j < 2*i+1; j++ {
			Writer.WriteString("*")
		}

		if i > 0 {
			Writer.WriteString("\n")
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
