package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 自分の得意な言語で
	// Let's チャレンジ！！
	//標準入力を受け取る
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	//受け取った値をスペース区切りで配列にする
	str := strings.Split(sc.Text(), " ")
	//int型の値を格納するスライスを宣言
	var sl []int
	//配列の中身をstr→intへ変換
	for _, v := range str {
		s, err := strconv.Atoi(v)
		if err != nil {
			break
		}
		sl = append(sl, s)
	}
	//変換した値を表示する
	for _, v := range sl {
		fmt.Println(v)
	}
}
