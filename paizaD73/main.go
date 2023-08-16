//問題文URL:https://paiza.jp/works/mondai/array_primer/array_primer__array_sort

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())
	var strArr []string = make([]string, num)
	sc.Scan()
	strArr = strings.Split(sc.Text(), " ")
	//strArrを数値型のスライスに変換
	var intArr []int
	for _, v := range strArr {
		i, _ := strconv.Atoi(v)
		intArr = append(intArr, i)
	}
	//intArrを昇順にソートする
	sort.Slice(intArr, func(i, j int) bool {
		return intArr[i] < intArr[j]
	})
	for _, v := range intArr {
		fmt.Println(v)
	}
}
