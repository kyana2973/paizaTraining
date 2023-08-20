// 問題文URL:https://paiza.jp/works/mondai/array_utilization_primer/array_utilization_primer__count
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
	counter := 0
	//入力される値に、Kと一致する値が存在したらカウントする
	for i := 0; i < N; i++ {
		sc.Scan()
		strArr = append(strArr, sc.Text())
		if strArr[i] == K {
			counter++
		}
	}
	fmt.Println(counter)
}
