//問題文URL:https://paiza.jp/works/mondai/array_primer/array_primer__array_reverse

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
	num, _ := strconv.Atoi(sc.Text())
	var strArr []string = make([]string, num)
	sc.Scan()
	strArr = strings.Split(sc.Text(), " ")
	//スライスを逆順で出力する
	for i := len(strArr) - 1; -1 < i; i-- {
		fmt.Println(strArr[i])
	}
}
