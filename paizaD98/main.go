// 問題文URL:https://paiza.jp/works/mondai/array_utilization_primer/array_utilization_primer__queue

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
	N, _ := strconv.Atoi(sc.Text())
	strArr := []string{}
	//N回の命令が入力される
	for i := 0; i < N; i++ {
		sc.Scan()
		order := sc.Text()
		//命令が"in"の場合は「in 10」とスペース区切りで整数が渡される
		//"out"の場合は、strArrの先頭にある値を削除する
		if strings.Contains(order, "in") {
			value := strings.Split(order, " ")
			strArr = append(strArr, value[1])
		} else if strings.Contains(order, "out") && len(strArr) > 0 {
			strArr = strArr[1:]
		}
	}
	//結果を改行区切りで出力する
	if len(strArr) > 0 {
		for _, v := range strArr {
			fmt.Println(v)
		}
	}
}
