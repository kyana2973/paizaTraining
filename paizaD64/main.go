//問題文URL:https://paiza.jp/works/mondai/string_primer/advance_step12

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
	B, _ := strconv.Atoi(sc.Text())

	//数列を逆順にする
	var rA []string
	for i := len(A) - 1; i > -1; i-- {
		rA = append(rA, A[i])
	}

	//AとBを掛け算する
	var sum int = 0
	var upNumFlg bool
	var countUp string
	var sumArr []string
	for i := 0; i < len(rA); i++ {
		a, _ := strconv.Atoi(rA[i])
		sum = a * B
		if upNumFlg {
			j, _ := strconv.Atoi(countUp)
			sum = sum + j
		}
		str := strconv.Itoa(sum)
		if len(str) > 1 {
			upNumFlg = true
			strArr := strings.Split(str, "")
			countUp = strArr[0]
			sumArr = append(sumArr, strArr[1])
		} else {
			upNumFlg = false
			sumArr = append(sumArr, str)
		}
	}

	if upNumFlg {
		sumArr = append(sumArr, countUp)
	}

	var output []string
	for i := len(sumArr) - 1; i > -1; i-- {
		output = append(output, sumArr[i])
	}

	fmt.Println(strings.Join(output, ""))
}
