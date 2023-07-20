//問題文URL:https://paiza.jp/works/mondai/logical_operation/logical_operation__basic_step9

package main

import (
	"fmt"
	"strings"
)

func main() {
	//0か1の整数がスペース区切りで２つ入力されるので、int型に変換してスライスに格納する
	var intArr []int = make([]int, 3)
	fmt.Scanf("%d %d %d", &intArr[0], &intArr[1], &intArr[2])

	//全加算器で出力する
	str := fmt.Sprintf("%02b", (intArr[0]+intArr[1]) + intArr[2])
	fmt.Println(strings.Join(strings.Split(str, ""), " "))
}
