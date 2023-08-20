// 問題文URL:https://paiza.jp/works/mondai/array_utilization_primer/array_utilization_primer__find

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
	num := strings.Split(sc.Text(), " ")
	N, _ := strconv.Atoi(num[0])
	K := num[1]
	strArr := []string{}
	checkFlg := false
	//スライスにKと一番最初に一致した順番を出力して、不一致の場合は-1を出力する
	for i := 0; i < N; i++ {
		sc.Scan()
		strArr = append(strArr, sc.Text())
		if strArr[i] == K {
			checkFlg = true
			fmt.Println(i + 1)
			break
		}
	}
	if !checkFlg {
		fmt.Println(-1)
	}
}
