//問題文URL:https://paiza.jp/works/mondai/array_primer/array_primer__2dmatrix_i-thoutput_step2

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
	strArr := strings.Split(sc.Text(), " ")
	K, _ := strconv.Atoi(strArr[0])
	L, _ := strconv.Atoi(strArr[1])
	//入力された3行の2次元スライスからK行、L列の値を出力する
	var outputArr [][]string
	for i := 0; i < 3; i++ {
		sc.Scan()
		v := strings.Split(sc.Text(), " ")
		outputArr = append(outputArr, v)
	}
	fmt.Println(outputArr[K-1][L-1])
}
