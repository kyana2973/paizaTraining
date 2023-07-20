//問題文URL:https://paiza.jp/works/mondai/logical_operation/logical_operation__basic_step4

package main

import (
	"fmt"
)

func main() {
	//0か1の整数がスペース区切りで２つ入力されるので、int型に変換してスライスに格納する
	var intArr []int = make([]int, 2)
	fmt.Scanf("%d %d", &intArr[0], &intArr[1])

	//排他的論理和を出力する
	fmt.Println(intArr[0] ^ intArr[1])
}
