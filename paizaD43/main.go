package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	//B = 判定用の文字列
	num := strings.Split(sc.Text(), " ")
	B := num[1]

	//判定対象のslice
	sc.Scan()
	numbers := strings.Split(sc.Text(), " ")

	//判定用の文字列が、sliceに存在するか判定する
	checkFlg := 0
	for _, v := range numbers {
		if v == B {
			checkFlg = 0
			break
		} else {
			checkFlg = 1
		}
	}

	//判定結果を出力する
	if checkFlg == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
