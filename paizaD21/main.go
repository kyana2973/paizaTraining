package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 1〜vまでの整数をスペース区切りで出力して改行する
func makeRow(v int) {
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
	//Mに格納する整数の数を格納
	N, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	//makeRowに渡す整数のスライス
	M := strings.Split(sc.Text(), " ")

	//N回だけmakeRow関数を実行する
	for i := 0; i < N; i++ {
		v, _ := strconv.Atoi(M[i])
		makeRow(v)
	}
}
