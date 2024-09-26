package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//모범답안.
//for n-2 {for n-1 { for n }} (On^3) 로 하는게 가장 깔끔하고 보기좋음.

func main() {
	//Input N -늘어놓을 카드의수 3 ~ 100 장 > Byte 3
	//Input M -최대 값이며 가까운값 찾기 > Byte 6
	//Input C [N]int - 늘어놓는 카드들 > int 8byte짜리 100 800bytes.   input buffer max 800
	//합쳐서 숫자가 가장크면서, 최대값 M을 넘지않는 인덱스 3개 찾기

	//ex) C[a], C[b], C[c] <= M
	//전략 1
	//두수의 합과 순서쌍을 담은 .. map[a+b][2]int.. 해시 만들기
	/*
	   a와 b 두가지값을 담아서 중복하지 않는 모든 경우의수를 찾아서 넣어야함 (순서쌍 2,4 또는 4,2 겹치면 뒤에오는건 탈락함)
	   중복값을 피하기위하여 순서쌍 검사 로직이 필요함 if-
	*/
	//
	Reader := bufio.NewReaderSize(os.Stdin, 800)
	Writer := bufio.NewWriter(os.Stdout)

	a, _, _ := Reader.ReadLine()
	nm := strings.Split(string(a), " ")
	N, _ := strconv.Atoi(nm[0])
	M, _ := strconv.Atoi(nm[1])

	c, _, _ := Reader.ReadLine()
	ca := strings.Split(string(c), " ")

	C := make([]int, N)
	//fmt.Println(N)
	//fmt.Println(C)
	for i := 0; i < N; i++ {
		//fmt.Println(ca[i])
		C[i], _ = strconv.Atoi(ca[i])
	}

	d := Twosum(C)

	//D에서 최대값 추출
	var max int = -1

	max1 := Twosumplus(d, C, max, M)
	Writer.WriteString(strconv.Itoa(max1))
	Writer.Flush()
}

//두정수의 합부터 모든 경우의수를 맵의 Key로서 넣고 순서쌍을 value로
//map[int(a+b)][2]int. 만약에 순서쌍이 (4,2) (2,4) 형태이면. 값이 같으니 체크하지않음.
//100장을 받으면 합의 경우의수는 100 x 99 x 98 이나옴.
//for n-2 {for n-1 { for n }} On^3 로 하는게 가장안정적임.
//a+b에서 100x99개를 구해야하지만, M값이 넘는 값을 탈락시켜야함.
//100x99-(M값을 넘는 수의 개수) 가 맵의 키로 들어가야함 A+B인 D로 치부함

//D의 값들에 C(D에서 사용한 인덱스를 제외한)을 더하여 배열을 구성한뒤에. 해당 구성한 배열에서 최대값을 추출하면 됨.
//

func Twosum(c []int) map[int][2]int {
	D := make(map[int][2]int)
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c); j++ {

			if i != j {
				D[c[i]+c[j]] = [2]int{i, j}

			}

		}
	}
	return D
}

func Twosumplus(c map[int][2]int, p []int, max int, m int) int {

	for k, v := range c {

		for i := 0; i < len(p); i++ {
			if !(i == v[0] || i == v[1]) {

				if k+p[i] <= m {

					if max < k+p[i] {
						max = k + p[i]
					}

				}
			}
		}
	}
	return max
}
