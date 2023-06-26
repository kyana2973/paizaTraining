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
	//最初の標準入力から行数を取得
	n, _ := strconv.Atoi(sc.Text())

	//出力用の２次元スライス
	var strSl [][]string
	//取得した行数分の標準入力を受け取って２次元スライスに格納する
	//入力される値は、「3 8 1 3」といった整数をスペースで区切った値となる
	for i := 0; i < n; i++ {
		sc.Scan()
		strArr := strings.Split(sc.Text(), " ")
		strSl = append(strSl, strArr)
	}

	//出力処理
	for i := 0; i < n; i++ {
		//二次元スライスの各行の先頭は、各行の値の数を表している
		//各行の値の数だけ出力したいので、その値を変数に格納する
		m, _ := strconv.Atoi(strSl[i][0])
		for j := 0; j < m; j++ {
			//各行の行末を出力した後は改行する
			if j == m-1 {
				fmt.Println(strSl[i][j+1])
				break
			}
			//先頭から値をスペース区切りで出力する
			fmt.Print(strSl[i][j+1] + " ")
		}
	}
}
