package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
라인 20줄
과목명 길이 최소 최대 1부터 50 까지
학점 1.2.3.4 . 0
등급은 A+, A0, B+,B0, C+, D+,F,P
at least this prob assure that there's no a case only P exist.

A+	4.5
A0	4.0
B+	3.5
B0	3.0
C+	2.5
C0	2.0
D+	1.5
D0	1.0
F	0.0

*/

func main() {
	var sub Sub
	var scoresum float64
	var eachscor float64
	sublist := make([]Sub, 20)
	for i := 0; i < 20; i++ {
		byt := Readline()
		sublist[i] = sub.SetScore(byt)

		//학점

		//eachscore, err := strconv.Atoi(sublist[i].Subscore)
		eachscore, _ := strconv.ParseFloat(sublist[i].Subscore, 64)
		if !(sublist[i].Subgrade == "P") {
			if sublist[i].Subgrade == "A+" {
				eachscor += eachscore * 4.5
			} else if sublist[i].Subgrade == "A0" {
				eachscor += eachscore * 4.0
			} else if sublist[i].Subgrade == "B+" {
				eachscor += eachscore * 3.5
			} else if sublist[i].Subgrade == "B0" {
				eachscor += eachscore * 3.0
			} else if sublist[i].Subgrade == "C+" {
				eachscor += eachscore * 2.5
			} else if sublist[i].Subgrade == "C0" {
				eachscor += eachscore * 2.0
			} else if sublist[i].Subgrade == "D+" {
				eachscor += eachscore * 1.5
			} else if sublist[i].Subgrade == "D0" {
				eachscor += eachscore * 1.0
			} else if sublist[i].Subgrade == "F" {
				eachscor += eachscore * 0
			}
			scoresum += eachscore
		}

	}
	avg := eachscor / scoresum

	fmt.Println(avg)
	//평균계산.

}

var (
	Reader *bufio.Reader
	Writer *bufio.Writer
)

func init() {
	Reader = bufio.NewReaderSize(os.Stdin, 50)
	Writer = bufio.NewWriter(os.Stdout)
}

func Readline() []byte {
	byt, _ := Reader.ReadBytes(13)
	//fmt.Println(byt)
	return byt
}

type Sub struct {
	Subname  string
	Subscore string
	Subgrade string
}

func (sub *Sub) SetScore(line []byte) Sub {
	//바이트 라인을 읽고 '32' 스페이스 만나면 스플릿 하고 버퍼 플러쉬 sub에 추가
	//지점저장
	//pointer

	counter := 0
	var tempslice []byte
	for _, byt := range line {

		if byt == 32 {

			if counter == 0 {
				sub.Subname = string(tempslice)
				tempslice = tempslice[:0]
			}

			if counter == 1 {
				sub.Subscore = string(tempslice)
				tempslice = tempslice[:0]
			}

			counter++
		} else if byt == 13 {
			sub.Subgrade = string(tempslice)
		} else {
			tempslice = append(tempslice, byt)
		}

	}
	return *sub
}
