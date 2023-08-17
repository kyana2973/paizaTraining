//問題文URL:https://paiza.jp/works/mondai/array_primer/array_primer__fibo

package main

import "fmt"

func main() {
	N := 0
	fmt.Scanf("%d", &N)
	//フィボナッチ数を出力する際に必要な初期値を設定する
	intArr := make([]int, N)
	intArr[0] = 0
	intArr[1] = 1
	//フィボナッチ数を出力する
	for i := 0; i < N; i++ {
		if i < 2 {
			fmt.Println(intArr[i])
		} else {
			intArr[i] = intArr[i-2] + intArr[i-1]
			fmt.Println(intArr[i])
		}
	}
}
