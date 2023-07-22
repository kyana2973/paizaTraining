package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str := sc.Text()
	//受け取った文字列が、paizaか確認する
	if str == "paiza" {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
