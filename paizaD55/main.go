package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str := sc.Text()
	sc.Scan()
	contain := sc.Text()

	//strの文字列中に、containの文字列が部分一致しているか判定する
	if strings.Contains(str, contain) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
