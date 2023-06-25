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

	//２回目から４回目の入力でスペース区切りの文字列が入力される
	//それを２次元スライスに格納する
	var strArr [][]string
	for i := 0; i < 3; i++ {
		sc.Scan()
		str := strings.Split(sc.Text(), " ")
		strArr = append(strArr, str)
	}

	//1行ずつint型に変換するための変数
	var chIntArr []int
	//変換された配列を格納するint型の２次元スライス
	var intArr [][]int
	for i := 0; i < len(strArr); i++ {
		//文字列の２次元スライスから1行ずつint型のスライスに変換する
		for j := 0; j < m; j++ {
			v, _ := strconv.Atoi(strArr[i][j])
			chIntArr = append(chIntArr, v)
		}
		//変換されたスライスをint型の２次元スライスに格納する
		intArr = append(intArr, chIntArr)
		//変換用のスライスを初期化
		chIntArr = nil
	}

	//int型の２次元スライスを行ごとに表示する
	// for i := 0; i < len(intArr); i++ {
	// 	fmt.Println(intArr[i])
	// }

	//int型の２次元スライスを１要素ずつ出力して、行末は改行して表示
	for i := 0; i < len(intArr); i++ {
		for j := 0; j < m; j++ {
			if j == m-1 {
				fmt.Println(intArr[i][j])
				break
			}
			fmt.Print(intArr[i][j])
		}
	}

}
