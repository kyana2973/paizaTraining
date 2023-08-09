package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	num := strings.Split(sc.Text(), " ")
	N, _ := strconv.Atoi(num[0])
	M := num[1]
	var strArr []string = make([]string, N)
	sc.Scan()
	strArr = strings.Split(sc.Text(), " ")
	var checkFlg bool = false
	//入力された文字列のスライスに、Mが含まれていたら"Yes"を、含まれていない場合は"No"を出力する
	for _, v := range strArr {
		if v == M {
			fmt.Println("Yes")
			checkFlg = true
			break
		}
	}
	if !checkFlg {
		fmt.Println("No")
	}
}
