package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//標準入力を受ける
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	//最初の値は行数を表す値になる
	rowCount, _ := strconv.Atoi(sc.Text())
	//N個の整数を格納するスライス
	var intArr []int

	//取得した行数分ループする
	for i := 0; i < rowCount; i++ {
		sc.Scan()
		//受け取った値を整数にしてスライスに格納
		v, _ := strconv.Atoi(sc.Text())
		intArr = append(intArr, v)
	}

	//スライスの中身を全て表示する
	for _, v := range intArr {
		fmt.Println(v)
	}
}
