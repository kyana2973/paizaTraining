package main

import (
	"bufio"
	"fmt"
	"os"
)

// 全体で使いたいのでグローバルに宣言
var sc = bufio.NewScanner(os.Stdin)

func main() {
	//10行入力したいので配列数10で宣言する
	var s [10]string

	//scanArr()によって返却された値を配列に格納
	for i := 0; i < len(s); i++ {
		s[i] = scanArr()
	}

	//s[i]を順番に出力する
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

// 標準入力を受けて値を返却
func scanArr() (s string) {
	sc.Scan()
	s = sc.Text()
	return
}
