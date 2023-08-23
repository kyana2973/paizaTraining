// 問題文URL:https://paiza.jp/works/mondai/array_utilization_primer/array_utilization_primer__exam

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
	//スコアを算出する対象の人数
	N, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	//点数に掛ける重み（５教科分）
	strWeight := strings.Split(sc.Text(), " ")
	scores := [][]string{}
	//各対象者のスコアを、文字列の二次元スライスに格納する
	for i := 0; i < N; i++ {
		sc.Scan()
		score := strings.Split(sc.Text(), " ")
		scores = append(scores, score)
	}
	results := [][]int{}
	//各教科に重みを掛けた値を、数値型の二次元スライスに格納する
	for i := 0; i < N; i++ {
		sum := []int{}
		for j := 0; j < 5; j++ {
			r, _ := strconv.Atoi(scores[i][j])
			w, _ := strconv.Atoi(strWeight[j])
			v := r * w
			sum = append(sum, v)
		}
		results = append(results, sum)
	}
	output := []int{}
	//各対象者の5教科合計を数値型スライスに格納する
	for i := 0; i < N; i++ {
		sum := 0
		for j := 0; j < 5; j++ {
			sum = sum + results[i][j]
		}
		output = append(output, sum)
	}
	//昇順にソートして、1番合計点が高い人の値を出力する
	sort.Ints(output)
	fmt.Println(output[len(output)-1])
}
