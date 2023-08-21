// 問題文URL:https://paiza.jp/works/mondai/array_utilization_primer/array_utilization_primer__reduce

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
	intArr := []int{}
	for i := 0; i < N; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		intArr = append(intArr, v)
	}
	//スライスの値を全てのペアで掛け算して出力する
	for i := 1; i < len(intArr); i++ {
		for j := 0; j < i; j++ {
			fmt.Println(intArr[j] * intArr[i])
		}
	}
}
