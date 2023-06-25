package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//標準入力取得
	sc := bufio.NewScanner(os.Stdin)
	//１回目は配列の列数が入力される
	sc.Scan()
	//列数をint型に変換
	m, _ := strconv.Atoi(sc.Text())

	//２回目から４回目まで入力が行われる
	//値は、整数がスペース区切りになった文字列が入力される → [1 2 3 4 5]といった文字列
	//２次元スライスにスペース区切りで格納する
	var strArr [][]string
	for i := 0; i < 3; i++ {
		sc.Scan()
		//値をスペース区切りで分解して配列としてstrに格納
		str := strings.Split(sc.Text(), " ")
		//２次元スライスに配列ごと格納
		strArr = append(strArr, str)
	}

	//出力処理
	for i := 0; i < len(strArr); i++ {
		//２次元スライスから1行ずつ出力
		for j := 0; j < m; j++ {
			//行末の値は改行して出力
			if j == m-1 {
				fmt.Println(strArr[i][j])
				break
			}
			//値と値の間にはスペースを含めて出力
			fmt.Print(strArr[i][j] + " ")
		}
	}

}
