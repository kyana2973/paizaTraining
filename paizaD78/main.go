//問題文URL:https://paiza.jp/works/mondai/array_primer/array_primer__deduplication

package main

import (
	"fmt"
	"sort"
)

func main() {
	intArr := []int{1, 3, 5, 1, 2, 3, 6, 6, 5, 1, 4}
	m := map[int]bool{}
	output := []int{}
	//intArrから重複した値を削除する
	for _, v := range intArr {
		if !m[v] {
			m[v] = true
			output = append(output, v)
		}
	}
	//昇順にする
	sort.Ints(output)
	//改行で全て出力する
	for _, v := range output {
		fmt.Println(v)
	}
}
