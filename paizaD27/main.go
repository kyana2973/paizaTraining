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
	//標準入力を受け取る
	strArr := strings.Split(sc.Text(), " ")
	//これから入力される回数
	N, _ := strconv.Atoi(strArr[0])
	//値をM桁で表示する際に使用する
	M := strArr[1]

	//N回の入力を受け取る
	var values []string = make([]string, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		values[i] = sc.Text()
	}

	//N個の値を全てM桁で表示する。元の値がM桁未満の場合は、左側にスペースを付加する
	for i := 0; i < N; i++ {
		fmt.Printf("%"+M+"s\n", values[i])
	}
}
