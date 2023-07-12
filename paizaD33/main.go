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
	//N、Mの値を取得する
	numbers := strings.Split(sc.Text(), " ")
	N, _ := strconv.Atoi(numbers[0])
	M, _ := strconv.Atoi(numbers[1])

	//整数の文字列がスペース区切りで入力されるので受け取る
	//スライスの要素数は、Nとする
	strArr := make([]string, N)
	sc.Scan()
	strArr = strings.Split(sc.Text(), " ")

	//スライスの中からM番目の値を出力する
	fmt.Println(strArr[M-1])

}
