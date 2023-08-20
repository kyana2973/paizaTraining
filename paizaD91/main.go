// 問題文URL:https://paiza.jp/works/mondai/array_utilization_primer/array_utilization_primer__insert

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
	N, _ := strconv.Atoi(sc.Text())
	strArr := []string{}
	for i := 0; i < N; i++ {
		sc.Scan()
		strArr = append(strArr, sc.Text())
	}
	sc.Scan()
	numValue := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(numValue[0])
	B := numValue[1]
	strArr = append(strArr[:n], strArr[n-1:]...)
	strArr[n] = B
	for _, v := range strArr {
		fmt.Println(v)
	}
}
