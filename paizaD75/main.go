//問題文URL:https://paiza.jp/works/mondai/array_primer/array_primer__elm_erase

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
	strArr := make([]string, N)
	sc.Scan()
	strArr = strings.Split(sc.Text(), " ")
	//スライスのM番目を削除して出力する
	strArr = append(strArr[:M-1], strArr[M:]...)
	for _, v := range strArr {
		fmt.Println(v)
	}
}
