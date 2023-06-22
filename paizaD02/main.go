package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//標準入力を受ける
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	//受けた値をスペース区切りで配列にする
	strArr := strings.Split(sc.Text(), " ")
	//配列を全て出力
	for i := 0; i < len(strArr); i++ {
		fmt.Println(strArr[i])
	}
}
