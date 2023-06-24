package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//最初の標準入力は要素数
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	count, _ := strconv.Atoi(sc.Text())

	//２つ目以降の標準入力は、文字列で小数点を表した値が入力される
	//それをfloat型に変換してスライスに格納する
	var floArr []float64
	for i := 0; i < count; i++ {
		sc.Scan()
		flo, _ := strconv.ParseFloat(sc.Text(), 64)
		floArr = append(floArr, flo)
	}

	//スライスの中身を全て表示する
	for _, v := range floArr {
		fmt.Println(v)
	}
}
