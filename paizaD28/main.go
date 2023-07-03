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
	//標準入力をスペース区切りで受け取る
	strArr := strings.Split(sc.Text(), " ")
	//Nは繰り返し回数
	N, _ := strconv.Atoi(strArr[0])
	//AとBは、表示する値
	A := strArr[1]
	B := strArr[2]

	//N回だけカンマスペース区切りで、AとBを出力する
	for i := 0; i < N; i++ {
		if i == N-1 {
			fmt.Println("(" + A + ", " + B + ")")
			break
		}
		fmt.Print("(" + A + ", " + B + "), ")
	}
}
