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
	//この後の入力回数を示す値を取得
	Q, _ := strconv.Atoi(sc.Text())

	//入力回数分ループする
	for i := 0; i < Q; i++ {
		sc.Scan()
		//入力を受け取って、スペース区切りでSplitして配列にする
		strArr := strings.Split(sc.Text(), " ")
		//Nには、出力する実数を格納
		N, _ := strconv.ParseFloat(strArr[0], 64)
		//Mには、表示する小数点以下の桁数を格納
		M := strArr[1]
		//Nを小数点第M位まで表示して、小数点第M位まで存在しない場合はゼロ埋めを行う
		fmt.Printf("%."+M+"f\n", N)
	}
}
