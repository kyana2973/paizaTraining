//問題文URL:https://paiza.jp/works/mondai/array_primer/array_primer__elm_insert

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
	M, _ := strconv.Atoi(num[1])
	K := num[2]
	strArr := make([]string, N)
	sc.Scan()
	strArr = strings.Split(sc.Text(), " ")
	//スライスのM番目にKを追加して出力する
	strArr = append(strArr[:M], strArr[M-1:]...)
	strArr[M-1] = K
	for _, v := range strArr {
		fmt.Println(v)
	}
}
