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
	//valuesの要素数が入力される
	N, _ := strconv.Atoi(sc.Text())
	values := make([]string, N)
	sc.Scan()
	//スペース区切りの整数が入力される
	values = strings.Split(sc.Text(), " ")
	sc.Scan()
	//numbersの要素数が入力される
	Q, _ := strconv.Atoi(sc.Text())
	numbers := make([]string, Q)
	sc.Scan()
	//スペース区切りの整数が入力される
	numbers = strings.Split(sc.Text(), " ")

	//valuesの中から、numbersの値番目だけ出力する
	for _, v := range numbers {
		i, _ := strconv.Atoi(v)
		fmt.Println(values[i-1])
	}
}
