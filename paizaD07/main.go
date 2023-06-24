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
	//標準入力をスペース区切りにして配列で受け取る
	strArr := strings.Split(sc.Text(), " ")
	//strArr[0]の値は、出力する要素数となる
	count, _ := strconv.Atoi(strArr[0])

	//出力する値をint型のスライスに格納
	var intArr []int
	//strArr[1]以降の値を全てint型に変換してスライスに格納
	for i := 1; i < count+1; i++ {
		value, _ := strconv.Atoi(strArr[i])
		intArr = append(intArr, value)
	}

	//スライスの中身を全て表示する
	for _, v := range intArr {
		fmt.Println(v)
	}
}
