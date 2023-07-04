package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var strArr []string
	//文字列がランダム回数入力されるので全てスライスに格納する
	for {
		sc.Scan()
		s := sc.Text()
		if len(s) == 0 {
			break
		}
		strArr = append(strArr, s)
	}

	//入力された文字列を連結する式を格納
	var output1 string
	//入力された文字列を全て連結した値を格納
	var output2 string
	//スライスの要素数分ループする
	for i := 0; i < len(strArr); i++ {
		//文字列を全て連結する
		output2 = output2 + strArr[i]
		if (i + 1) == len(strArr) {
			//最後の文字列の場合は、＋ではなく＝を付与する
			output1 = output1 + strArr[i] + " = "
			//式 = 値という形式で出力する
			fmt.Println(output1 + output2)
		} else {
			//文字列を連結する式を作成
			output1 = output1 + strArr[i] + " + "
		}
	}

}
