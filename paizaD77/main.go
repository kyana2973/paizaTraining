//問題文URL:https://paiza.jp/works/mondai/array_primer/array_primer__string_dictionary_boss

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	num := strings.Split(sc.Text(), " ")
	N, _ := strconv.Atoi(num[0])
	K, _ := strconv.Atoi(num[1])
	strArr := make([]string, N)
	sc.Scan()
	strArr = strings.Split(sc.Text(), " ")
	//文字列スライスを昇順にする
	sort.Strings(strArr)
	//K番目の値を出力する
	fmt.Println(strArr[K-1])
}
