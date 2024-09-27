////Reference code - Baekjoon no.1 code. c++ code changed chanaged to Go.
//// not mine.

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

//내 코드
/*
func main() {
	defer writer.Flush()

	var basket []int

	input, _ := io.ReadAll(reader)
	commands := strings.Split(string(input), "\n")


	N, _ := strconv.Atoi(commands[0])


	for i := 1; i <= N; i++ {
		command := commands[i]
		if len(command) > 1 && command[0] == '1' {

			x, _ := strconv.Atoi(command[2:])
			Stack(&basket, x)
		} else {

			ms, _ := strconv.Atoi(command)
			if ms == 5 {
				if Isempty(&basket) {
					writer.WriteString("-1\n")
				} else {
					writer.WriteString(strconv.Itoa(basket[len(basket)-1]) + "\n")
				}
			} else if ms == 3 {
				writer.WriteString(strconv.Itoa(CheckInt(&basket)) + "\n")
			} else if ms == 4 {
				writer.WriteString(strconv.Itoa(Isemp(&basket)) + "\n")
			} else if ms == 2 {
				writer.WriteString(strconv.Itoa(Pop(&basket)) + "\n")
			}
		}
	}
}

var (
	reader *bufio.Reader
	writer *bufio.Writer
)

func init() {
	reader = bufio.NewReaderSize(os.Stdin, 9000000)
	writer = bufio.NewWriter(os.Stdout)
}

func Stack(basket *[]int, thing int) *[]int {
	*basket = append(*basket, thing)
	return basket
}

func Pop(basket *[]int) (thing int) {
	if Isempty(basket) {
		return -1
	}
	top := len(*basket) - 1
	thing = (*basket)[top]
	*basket = (*basket)[:top] //Remove Top side
	return thing
}

func Isempty(basket *[]int) bool {
	return len(*basket) == 0
}

func Isemp(basket *[]int) int {
	if len(*basket) == 0 {
		return 1
	} else {
		return 0
	}
}

func CheckInt(basket *[]int) int {
	return len(*basket)
}
*/
