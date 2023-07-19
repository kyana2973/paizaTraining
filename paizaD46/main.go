//問題文URL:https://paiza.jp/works/mondai/data_structure/data_structure__set_boss

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
	//numは、入力される値の要素数となる
	num, _ := strconv.Atoi(sc.Text())
	var numbers []string = make([]string, num*2)
	sc.Scan()
	//数列Aが入力されるので、スペースで区切ってからスライスに格納する
	numbers = strings.Split(sc.Text(), " ")
	sc.Scan()
	//数列Bが入力されるので、スペースで区切ってからスライスに格納する
	numbers = append(numbers, strings.Split(sc.Text(), " ")...)

	//スライスの中身で、重複している値は除外して出力用のスライスへ、int型に変換してから格納する
	valSort := make(map[string]bool)
	outputInt := []int{}
	for _, v := range numbers {
		if !valSort[v] {
			valSort[v] = true
			i, _ := strconv.Atoi(v)
			outputInt = append(outputInt, i)
		}
	}

	//値を昇順でソートする
	sort.Ints(outputInt)

	//スライスの中身をstring型に変換する
	outputStr := []string{}
	for _, v := range outputInt {
		s := strconv.Itoa(v)
		outputStr = append(outputStr, s)
	}

	//スペース区切りで出力する
	fmt.Println(strings.Join(outputStr, " "))
}
