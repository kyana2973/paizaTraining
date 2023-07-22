//問題文URL:https://paiza.jp/works/mondai/logical_operation/logical_operation__basic_boss

package main

import (
	"fmt"
)

func main() {
	//0か1の整数がスペース区切りで４つ入力されるので、int型に変換してスライスに格納する
	var intArr []int = make([]int, 4)
	fmt.Scanf("%d %d %d %d", &intArr[0], &intArr[1], &intArr[2], &intArr[3])

	//NOT((NOT(intArr[0]) AND NOT(intArr[1])) OR NOT(intArr[2])) XOR intArr[3] を求める
	//ド・モルガンの法則を適用して以下式で出力
	fmt.Println(((intArr[0] | intArr[1]) & intArr[2]) ^ intArr[3])
}
