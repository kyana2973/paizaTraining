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
	//スペース区切りの整数が文字列で入力される
	intArr := strings.Split(sc.Text(), " ")

	//minとmaxを宣言しておく
	var max int = 0
	var min int = 0

	//intArrの中から最大値と最小値を取り出して、最大値ー最小値を算出して出力する
	for i, v := range intArr {
		n, _ := strconv.Atoi(v)
		if i == 0 {
			min = n
		} else if min > n {
			min = n
		}

		if max < n {
			max = n
		}
	}

	fmt.Println(max - min)
}
