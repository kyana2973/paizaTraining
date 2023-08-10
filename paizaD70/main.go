//問題文URL:https://paiza.jp/works/mondai/array_primer/array_primer__elm_rewrite

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
	A := num[0]
	B := num[1]
	N, _ := strconv.Atoi(num[2])
	var strArr []string = make([]string, N)
	sc.Scan()
	strArr = strings.Split(sc.Text(), " ")
	//スライスの中にAが存在したら、Bに変換する
	for i, v := range strArr {
		if v == A {
			strArr[i] = B
		}
		fmt.Println(strArr[i])
	}
}
