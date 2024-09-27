package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var (
	Reader *bufio.Reader
	Writer *bufio.Writer
)

func init() {
	Reader = bufio.NewReaderSize(os.Stdin, 1)
	Writer = bufio.NewWriter(os.Stdout)
}

func Readstr() string {
	str, _, _ := Reader.ReadLine()
	return string(str)
}

// need to exclude 10/13 and
func Readbyt() []byte {
	bt, _ := Reader.ReadBytes(13)
	renewdbt := bt[:len(bt)-1]
	return renewdbt
}

func Reversing(str []byte) []byte {

	reversingstr := make([]byte, 0)
	for i := len(str) - 1; i >= 0; i-- {
		reversingstr = append(reversingstr, str[i])
	}
	fmt.Println(str)
	fmt.Println(reversingstr)
	return reversingstr
}

func Validating(ab, ba []byte) bool {

	return bytes.Equal(ab, ba)
}
func main() {
	defer Writer.Flush()
	word := Readbyt()
	reword := Reversing(word)
	if Validating(word, reword) {
		Writer.WriteByte('1')
	} else {
		Writer.WriteByte('0')
	}
}
