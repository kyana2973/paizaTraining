package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str := sc.Text()

	var output string
	//受け取った文字列を、1文字ずつ小文字か大文字かを判定して、小文字<=>大文字を変換する
	for _, v := range str {
		if unicode.IsUpper(v) {
			output = output + strings.ToLower(string(v))
		} else if unicode.IsLower(v) {
			output = output + strings.ToUpper(string(v))
		}
	}

	fmt.Println(output)
}
