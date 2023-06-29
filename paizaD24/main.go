package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	//標準入力から整数の値を受け取る
	str := sc.Text()
	//出力桁が3桁になるようにスペースを左側に付加する
	fmt.Printf("%3s\n", str)
}
