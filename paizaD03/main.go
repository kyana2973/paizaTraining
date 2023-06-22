package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//スキャンを宣言
	sc := bufio.NewScanner(os.Stdin)
	//入力値を格納するスライス
	var sl []int

	for {
		//入力を受け取る
		sc.Scan()
		//受け取った値をint型にして一時退避
		i, err := strconv.Atoi(sc.Text())
		//入力値が無い場合はループを抜ける
		if err != nil {
			break
		}
		//値をスライスに格納する
		sl = append(sl, i)
	}

	//スライス内を全て出力する
	for _, v := range sl {
		fmt.Println(sl[v])
	}
}
