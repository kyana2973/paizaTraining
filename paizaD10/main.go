package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//最初の入力値は要素数を表しているためcountに格納する
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	count, _ := strconv.Atoi(sc.Text())

	//８回目の入力を変数に格納する
	var str string
	for i := 0; i < count; i++ {
		sc.Scan()
		if i == 7 {
			str = sc.Text()
		}
	}

	//８回目の入力のみを表示する
	fmt.Println(str)
}
