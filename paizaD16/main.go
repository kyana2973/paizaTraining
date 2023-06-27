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
	//複数桁の整数が標準入力される
	str := sc.Text()

	//スライスに1文字ずつ格納していく
	var strArr []string = strings.Split(str, "")
	//文字列数を３で割って余りを格納
	var overLen int = len(str) % 3

	//3桁ずつカンマを入れて出力し改行する
	//一番右の小さい位から3桁ずつにカンマが入るようにする
	//「12,345,678」のように右から3桁ずつカンマが入っているように見せる必要がある
	for i, s := range strArr {
		//余りがゼロの場合は、i=0の時に先頭にカンマが付いてしまうので、i=0は除外する
		if i%3 == overLen && i != 0 {
			fmt.Print(",")
		}
		fmt.Print(s)
	}
	fmt.Println()
}
