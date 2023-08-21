// 問題文URL:https://paiza.jp/works/mondai/array_utilization_primer/array_utilization_primer__delete

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	n, _ := strconv.Atoi(sc.Text())
	//スライスのn番目を削除して出力する
	strArr = append(strArr[:n-1], strArr[n:]...)
	for _, v := range strArr {
		fmt.Println(v)
	}
}
