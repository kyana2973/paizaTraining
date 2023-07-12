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
	//N = 最初の要素数、Q = 関数で操作する回数
	numOfOpe := strings.Split(sc.Text(), " ")
	N, _ := strconv.Atoi(numOfOpe[0])
	Q, _ := strconv.Atoi(numOfOpe[1])

	//N個の要素を受け取る
	values := make([]string, N)
	sc.Scan()
	values = strings.Split(sc.Text(), " ")

	//関数の番号
	var opeNum string
	//スライスに値を追加する際の変数
	var pushNum string
	for i := 0; i < Q; i++ {
		sc.Scan()
		s := sc.Text()
		if len(s) > 1 {
			strArr := strings.Split(s, " ")
			opeNum = strArr[0]
			pushNum = strArr[1]
		} else if len(s) == 1 {
			opeNum = s
		}

		switch opeNum {
		case "0":
			//スライスの末尾に値を追加する
			values = append(values, pushNum)
		case "1":
			//スライスの末尾から値を1個削除する
			values = values[:len(values)-1]
		case "2":
			//現時点でのスライスの要素を全てスペース区切りで出力する
			for i, v := range values {
				if i == len(values)-1 {
					fmt.Println(v)
				} else {
					fmt.Print(v + " ")
				}
			}
		}
	}
}
