package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	//整数9桁の標準入力を受ける
	str := sc.Text()

	//3桁ごとにカンマを付与して改行する
	for i := 0; i < len(str); i++ {
		//改行処理
		if i == len(str)-1 {
			fmt.Println(str[i-2:])
			break
		}

		//3桁ごとにカンマを付与する
		if (i+1)%3 == 0 {
			fmt.Print(str[i-2:i+1] + ",")
		}
	}
}
