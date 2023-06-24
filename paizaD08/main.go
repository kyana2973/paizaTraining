package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//最初の標準入力を受け取る、この値は出力する要素数になる
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	count, _ := strconv.Atoi(sc.Text())

	//2個目の標準入力はスペース区切り整数が連なっている文字列で、スペース区切りでスライスに格納
	sc.Scan()
	var strArr []string = strings.Split(sc.Text(), " ")

	//int型のスライスにstrArrをint型に変換してから全て格納する
	var intArr []int
	for i := 0; i < count; i++ {
		chInt, _ := strconv.Atoi(strArr[i])
		intArr = append(intArr, chInt)
	}

	//intArrを全て表示する
	for _, v := range intArr {
		fmt.Println(v)
	}
}
