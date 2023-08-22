// 問題文URL:https://paiza.jp/works/mondai/array_utilization_primer/array_utilization_primer__unique

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
	m := map[string]bool{}
	//入力される値が重複している場合は、削除してスライスに格納し出力する
	for i := 0; i < N; i++ {
		sc.Scan()
		v := sc.Text()
		if !m[v] {
			m[v] = true
			strArr = append(strArr, v)
		}
	}
	for _, v := range strArr {
		fmt.Println(v)
	}
}
