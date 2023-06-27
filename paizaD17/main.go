package main

import (
	"fmt"
	"strconv"
)

func main() {
	var kuku []string
	//九九表を作成
	//スライスに九九の答えを全て格納
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			v := i * j
			s := strconv.Itoa(v)
			kuku = append(kuku, s)
		}
	}

	//9行で改行するために余りを格納
	overLen := len(kuku) % 9

	//9行で改行して値と値の間にはスペースを付加して出力する
	for i := 0; i < len(kuku); i++ {
		if (i+1)%9 == overLen {
			fmt.Println(kuku[i])
		} else {
			fmt.Print(kuku[i] + " ")
		}
	}
}
