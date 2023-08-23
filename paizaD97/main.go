// 問題文URL:https://paiza.jp/works/mondai/array_utilization_primer/array_utilization_primer__offer

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
	num := strings.Split(sc.Text(), " ")
	//試験を受けた人数
	N, _ := strconv.Atoi(num[0])
	//合格点
	K, _ := strconv.Atoi(num[1])
	//試験後に辞退した人数（順位が高い順から辞退した人数）
	M, _ := strconv.Atoi(num[2])
	scores := []int{}
	//合格点以上の人だけスライスに格納する
	for i := 0; i < N; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		if v >= K {
			scores = append(scores, v)
		}
	}
	//辞退した人数より合格点以上の人が多かったら、合格した人の人数を出力する
	if len(scores)-M >= 0 {
		sort.Ints(scores)
		count := []int{}
		count = append(count, scores[:len(scores)-M]...)
		fmt.Println(len(count))
	} else {
		//辞退した人数の方が多い場合は、合格者はいないのでゼロとなる
		fmt.Println(0)
	}
}
