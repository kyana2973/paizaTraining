//問題文URL:https://paiza.jp/works/mondai/array_primer/array_primer__array_elm_change

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
	m := map[string]int{}
	m["A"], _ = strconv.Atoi(num[0])
	m["B"], _ = strconv.Atoi(num[1])
	m["N"], _ = strconv.Atoi(num[2])
	var strArr []string = make([]string, m["N"])
	sc.Scan()
	strArr = strings.Split(sc.Text(), " ")
	A := strArr[m["A"]-1]
	B := strArr[m["B"]-1]
	for i := range strArr {
		if i == m["A"]-1 {
			strArr[i] = B
		} else if i == m["B"]-1 {
			strArr[i] = A
		}
	}
	for _, v := range strArr {
		fmt.Println(v)
	}
}
