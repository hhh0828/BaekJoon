package main

import (
	"os"
	"strconv"
)

const SIZE = 1 << 20

var (
	arRBuffer   [SIZE]byte
	arWBuffer   [SIZE]byte
	nReadIndex  = SIZE
	nWriteIndex int
)

func ReadChar() byte {
	if nReadIndex == SIZE {
		_, err := os.Stdin.Read(arRBuffer[:])
		if err != nil {
			panic(err)
		}
		nReadIndex = 0
	}
	c := arRBuffer[nReadIndex]
	nReadIndex++
	return c
}

func ReadInt() int {
	cRead := ReadChar()
	for (cRead < '0' || cRead > '9') && cRead != '-' {
		cRead = ReadChar()
	}
	nRet := 0
	bNeg := cRead == '-'
	if bNeg {
		cRead = ReadChar()
	}
	for cRead >= '0' && cRead <= '9' {
		nRet = nRet*10 + int(cRead-'0')
		cRead = ReadChar()
	}
	if bNeg {
		return -nRet
	}
	return nRet
}

func WriteChar(cWrite byte) {
	if nWriteIndex >= SIZE {
		Flush()
	}
	arWBuffer[nWriteIndex] = cWrite
	nWriteIndex++
}

func WriteInt(nWrite int) {
	str := strconv.Itoa(nWrite)
	length := len(str)
	if nWriteIndex+length+1 >= SIZE {
		Flush()
	}
	copy(arWBuffer[nWriteIndex:], str)
	nWriteIndex += length
	WriteChar('\n')
}

func Flush() {
	_, err := os.Stdout.Write(arWBuffer[:nWriteIndex])
	if err != nil {
		panic(err)
	}
	nWriteIndex = 0
}

func main() {
	N := ReadInt()
	stack := make([]int, 0, 1000001)
	for i := 0; i < N; i++ {
		cmd := ReadInt()
		switch cmd {
		case 1:
			X := ReadInt()
			stack = append(stack, X)
		case 2:
			if len(stack) > 0 {
				WriteInt(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			} else {
				WriteInt(-1)
			}
		case 3:
			WriteInt(len(stack))
		case 4:
			if len(stack) == 0 {
				WriteInt(1)
			} else {
				WriteInt(0)
			}
		case 5:
			if len(stack) > 0 {
				WriteInt(stack[len(stack)-1])
			} else {
				WriteInt(-1)
			}
		}
	}
	Flush()
}
