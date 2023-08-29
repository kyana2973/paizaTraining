// 問題文URL:https://paiza.jp/works/mondai/array_utilization_primer/array_utilization_primer__find_pair

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())
	heightArr := []int{}
	//全ての身長データを受け取る
	for i := 0; i < N; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		heightArr = append(heightArr, v)
	}
	//昇順にする
	sort.Ints(heightArr)
	pair := make([]int, 2)
	compare := 0
	//一番身長差が無いペアを選定する
	for i := 1; i < len(heightArr)-1; i++ {
		//最初の2人の身長と身長差を初期設定する
		if i == 1 {
			pair[0] = heightArr[0]
			pair[1] = heightArr[1]
			//昇順で身長差を出すとマイナス値になるので絶対値に変換して格納する
			compare = int(math.Abs(float64(pair[0]) - float64(pair[1])))
		}
		//昇順で身長差を出すとマイナス値になるので絶対値に変換して格納する
		v := int(math.Abs(float64(heightArr[i]) - float64(heightArr[i+1])))
		//身長差がcompare変数より低かったら更新する
		if compare > v {
			pair[0] = heightArr[i]
			pair[1] = heightArr[i+1]
		}
	}
	//ペアの身長を昇順の改行区切りで出力する
	for i := 0; i < len(pair); i++ {
		fmt.Println(pair[i])
	}
}
