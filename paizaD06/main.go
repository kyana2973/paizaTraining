package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//標準入力を受ける
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	//最初の値は出力する値の数
	count, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	//２つ目の値は、スペース区切りで連なっている整数の値
	//それをスペース区切りで分けて配列に格納する
	intArr := strings.Split(sc.Text(), " ")

	//countの値分の行数を表示する
	for i := 0; i < count; i++ {
		fmt.Println(intArr[i])
	}
}
