package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	reg := "[/:]"
	sc.Scan()
	//日付のフォーマットで入力される文字列を、「/ or :」で分割する
	strArr := regexp.MustCompile(reg).Split(sc.Text(), -1)
	for _, v := range strArr {
		fmt.Println(v)
	}
}
