// 問題文URL:https://paiza.jp/works/mondai/array_utilization_primer/array_utilization_primer__exist
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
	K := num[1]
	strArr := []string{}
	checkFlg := false
	//スライスにKと一致する値が存在したら"Yes"を、不一致の場合は"No"を出力する
	for i := 0; i < N; i++ {
		sc.Scan()
		strArr = append(strArr, sc.Text())
		if strArr[i] == K {
			checkFlg = true
		}
	}
	if checkFlg {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
