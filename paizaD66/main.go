//問題文URL:https://paiza.jp/works/mondai/array_primer/array_primer__2dmatrix_input_step4

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
	//N行、M列の値を改行して出力する
	var strArr []string = make([]string, M)
	var allStrArr [][]string
	for i := 0; i < N; i++ {
		sc.Scan()
		strArr = strings.Split(sc.Text(), " ")
		allStrArr = append(allStrArr, strArr)
	}
	for i := 0; i < N; i++ {
		fmt.Println(strings.Join(allStrArr[i], " "))
	}
}
