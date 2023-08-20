// 問題文URL:https://paiza.jp/works/mondai/array_utilization_primer/array_utilization_primer__max

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())
	intArr := []int{}
	//入力される値を数値型スライスに格納して、その中から最大値を出力する
	for i := 0; i < N; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		intArr = append(intArr, v)
	}
	sort.Ints(intArr)
	fmt.Println(intArr[len(intArr)-1])
}
