// 問題文URL:https://paiza.jp/works/mondai/array_utilization_primer/array_utilization_primer__swap

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
	N, _ := strconv.Atoi(sc.Text())
	strArr := []string{}
	for i := 0; i < N; i++ {
		sc.Scan()
		strArr = append(strArr, sc.Text())
	}
	sc.Scan()
	num := strings.Split(sc.Text(), " ")
	x, _ := strconv.Atoi(num[0])
	y, _ := strconv.Atoi(num[1])
	//スライスのx,y番目の値を入れ替えて、改行区切りで出力する
	strArr[x-1], strArr[y-1] = strArr[y-1], strArr[x-1]
	for _, v := range strArr {
		fmt.Println(v)
	}
}
