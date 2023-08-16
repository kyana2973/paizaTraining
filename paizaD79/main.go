//問題文URL:https://paiza.jp/works/mondai/array_primer/array_primer__array_minmax

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
	N, _ := strconv.Atoi(sc.Text())
	strArr := make([]string, N)
	sc.Scan()
	strArr = strings.Split(sc.Text(), " ")
	intArr := []int{}
	for _, v := range strArr {
		i, _ := strconv.Atoi(v)
		intArr = append(intArr, i)
	}
	sort.Ints(intArr)
	//最大値と最小値をスペース区切りで出力する
	fmt.Println(strconv.Itoa(intArr[len(intArr)-1]) + " " + strconv.Itoa(intArr[0]))
}
