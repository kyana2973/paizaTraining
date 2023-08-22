// 問題文URL:https://paiza.jp/works/mondai/array_utilization_primer/array_utilization_primer__resize

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
	n, _ := strconv.Atoi(num[1])
	strArr := make([]string, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		strArr = append(strArr, sc.Text())
	}
	editArr := make([]string, n)
	//要素数がn個のスライスに、要素数がN個のスライスからn番目までコピーする
	//要素数の関係が、N < nだった場合はゼロ埋めする
	for i := 0; i < n; i++ {
		if i > N-1 {
			editArr = append(editArr, "0")
		} else {
			editArr = append(editArr, strArr[i])
		}
	}
	for _, v := range editArr {
		fmt.Println(v)
	}
}
