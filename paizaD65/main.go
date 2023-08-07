//問題文URL:https://paiza.jp/works/mondai/array_primer/array_primer__2dmatrix_input_step2

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
	M, _ := strconv.Atoi(sc.Text())
	var strArr []string = make([]string, M)
	var allStrArr [][]string
	for i := 0; i < 5; i++ {
		sc.Scan()
		strArr = strings.Split(sc.Text(), " ")
		allStrArr = append(allStrArr, strArr)
	}
	for i := 0; i < len(allStrArr); i++ {
		fmt.Println(strings.Join(allStrArr[i], " "))
	}
}
