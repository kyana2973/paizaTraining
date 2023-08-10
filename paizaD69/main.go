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
	num := strings.Split(sc.Text(), " ")
	N := num[0]
	M, _ := strconv.Atoi(num[1])
	strArr := make([]string, M)
	sc.Scan()
	strArr = strings.Split(sc.Text(), " ")
	//入力された文字列スライスの中からNと一致する文字列が存在したら、その文字列が左から何番目かを出力する
	for i, v := range strArr {
		if v == N {
			fmt.Println(i + 1)
		}
	}
}
