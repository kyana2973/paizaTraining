// 問題文URL:https://paiza.jp/works/mondai/array_utilization_primer/array_utilization_primer__reverse

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
	//スライスを反転して改行区切りで出力する
	reverseArray(strArr)
	for _, v := range strArr {
		fmt.Println(v)
	}
}

func reverseArray(arr []string) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
