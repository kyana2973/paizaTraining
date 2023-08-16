//問題文URL:https://paiza.jp/works/mondai/array_primer/array_primer__array_ave_step3

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
	strArr := make([]string, N)
	sc.Scan()
	strArr = strings.Split(sc.Text(), " ")
	sum := 0
	intArr := []int{}
	for _, v := range strArr {
		i, _ := strconv.Atoi(v)
		intArr = append(intArr, i)
		sum += i
	}
	avg := float64(sum) / float64(len(strArr))
	//intArrの平均値以上の値のみを出力する
	for _, v := range intArr {
		if float64(v) >= avg {
			fmt.Println(v)
		}
	}
}
