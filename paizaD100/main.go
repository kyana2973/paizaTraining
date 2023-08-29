// 問題文URL:https://paiza.jp/works/mondai/array_utilization_primer/array_utilization_primer__line_up

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
	sheets, _ := strconv.Atoi(num[1])
	deletes, _ := strconv.Atoi(num[2])
	sheetArr := []string{}
	//全シート番号を順番に受け取る
	for i := 0; i < sheets; i++ {
		sc.Scan()
		sheetArr = append(sheetArr, sc.Text())
	}
	delArr := []string{}
	//sheetArrの先頭からdeletes番目まで削除してdelArrに格納する
	delArr = append(delArr, sheetArr[deletes:]...)
	m := map[string]bool{}
	output := []string{}
	//残ったシート番号をスライスの中から先頭を優先にして、重複している値を除いてoutputに格納する。
	for _, v := range delArr {
		if !m[v] {
			m[v] = true
			output = append(output, v)
		}
	}
	//改行区切りで出力する
	for _, v := range output {
		fmt.Println(v)
	}
}
