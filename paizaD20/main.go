package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 1〜vまでの値を1行で出力する。その際にスペース区切りで行末は改行する
func triangle(v int) {
	for i := 0; i < v; i++ {
		s := strconv.Itoa(i + 1)
		if (i + 1) == v {
			fmt.Println(s)
			break
		}
		fmt.Print(s + " ")
	}
}
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	//行数となる値を受け取る
	N, _ := strconv.Atoi(sc.Text())

	//行数分だけtriangle関数を実行する
	for i := 0; i < N; i++ {
		triangle(i + 1)
	}
}
