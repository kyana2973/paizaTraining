//問題文URL:https://paiza.jp/works/mondai/string_primer/advance_step11

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	A := strings.Split(sc.Text(), "")
	sc.Scan()
	B := strings.Split(sc.Text(), "")

	//数列を逆順にする
	var rA []string
	var rB []string
	for i := len(A) - 1; i > -1; i-- {
		rA = append(rA, A[i])
		rB = append(rB, B[i])
	}

	//AとBを足し算する
	var sum int = 0
	var upNumFlg bool
	var sumArr []string = make([]string, len(A))
	for i := 0; i < len(rA); i++ {
		a, _ := strconv.Atoi(rA[i])
		b, _ := strconv.Atoi(rB[i])
		sum = a + b
		//前回の計算で繰り上げが発生していたら「+1」する
		if upNumFlg {
			sum = sum + 1
		}
		str := strconv.Itoa(sum)
		if len(str) > 1 {
			strArr := strings.Split(str, "")
			upNumFlg = true
			sumArr = append(sumArr, strArr[1])
		} else {
			upNumFlg = false
			sumArr = append(sumArr, str)
		}
	}

	var output []string
	//最後の計算で繰り上げが起きていたら、「1」をスライスに追加する
	if upNumFlg {
		output = append(output, "1")
	}
	//正規の順に戻す
	for i := len(sumArr) - 1; i > -1; i-- {
		output = append(output, sumArr[i])
	}

	fmt.Println(strings.Join(output, ""))
}
