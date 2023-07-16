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
	//N = 文字列リストの数、Q = 照合リストの数
	N, _ := strconv.Atoi(num[0])
	Q, _ := strconv.Atoi(num[1])

	//入力される文字列リストを、sliceに格納する
	strArr := make([]string, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		strArr[i] = sc.Text()
	}

	//文字列リスト = 照合リストとなった場合は、一致した時の文字列リストの行数を出力する
	//一致しなかったものは、「-1」と出力する
	checkFlg := 0
	for i := 0; i < Q; i++ {
		sc.Scan()
		for i, v := range strArr {
			if v == sc.Text() {
				checkFlg = 0
				fmt.Println(i + 1)
				break
			} else {
				checkFlg = 1
			}
		}
		if checkFlg == 1 {
			fmt.Println(-1)
		}
	}
}
