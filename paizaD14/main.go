package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var strArr [10]string
	//標準入力を10回受ける
	for i := 0; i < len(strArr); i++ {
		sc.Scan()
		strArr[i] = sc.Text()
	}

	//配列の文字列を全て連結し、文字間にスペースを付加して行末で改行する
	for i, s := range strArr {
		if i == len(strArr)-1 {
			fmt.Println(s)
			break
		}
		fmt.Print(s + " ")
	}
}
