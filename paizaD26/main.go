package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	//標準入力で数値を受け取る
	i, _ := strconv.Atoi(sc.Text())
	//数値が3桁になるよに不足分は、左側からゼロ埋めをする
	fmt.Printf("%03d\n", i)
}
