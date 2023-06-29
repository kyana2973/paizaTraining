package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	//出力する値の数、1行に出力する値の数を取得
	index := strings.Split(sc.Text(), " ")
	N, _ := strconv.Atoi(index[0])
	M, _ := strconv.Atoi(index[1])

	//出力する値をスライスに格納
	var A []string = make([]string, N)
	sc.Scan()
	A = strings.Split(sc.Text(), " ")

	//1行に出力する値の数を示す数値をスライスに格納
	var B []string = make([]string, M)
	sc.Scan()
	B = strings.Split(sc.Text(), " ")

	//出力すインデックスを保存しておく変数
	saveIndex := 0
	for i := 0; i < M; i++ {
		//出力する値の数を数値型で格納
		count, _ := strconv.Atoi(B[i])
		for j := 0; j < count; j++ {
			if (j + 1) == count {
				fmt.Println(A[saveIndex])
				saveIndex++
				break
			}
			fmt.Print(A[saveIndex] + " ")
			saveIndex++
		}
	}
}
