// 問題文URL:https://paiza.jp/works/mondai/array_utilization_primer/array_utilization_primer__bowling

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
	N, _ := strconv.Atoi(num[0])
	Q, _ := strconv.Atoi(num[1])
	humans := []string{}
	//昇順で整数を格納する
	for i := 0; i < N; i++ {
		humans = append(humans, strconv.Itoa(i+1))
	}
	for i := 0; i < Q; i++ {
		sc.Scan()
		//命令と引数を受け取る
		order := strings.Split(sc.Text(), " ")
		switch order[0] {
		case "swap":
			humans = swap(humans, order[1], order[2])
			break
		case "reverse":
			humans = reverse(humans)
			break
		case "resize":
			humans = resize(humans, order[1])
			break
		}
	}
	//最終結果を出力する
	for _, v := range humans {
		fmt.Println(v)
	}
}

// humansのA番目とB番目を入れ替える
func swap(humans []string, A, B string) []string {
	aI, _ := strconv.Atoi(A)
	bI, _ := strconv.Atoi(B)
	aVal := humans[aI-1]
	bVal := humans[bI-1]
	humans[aI-1] = bVal
	humans[bI-1] = aVal
	return humans
}

// humansの並びを反転させる
func reverse(humans []string) []string {
	for i, j := 0, len(humans)-1; i < j; i, j = i+1, j-1 {
		humans[i], humans[j] = humans[j], humans[i]
	}
	return humans
}

// humansの長さを、先頭からC番目までに加工して残りは切り捨てる
// humansの要素数が、C以下の場合は処理を行わない
func resize(humans []string, C string) []string {
	c, _ := strconv.Atoi(C)
	if len(humans) > c {
		humans = append(humans[:0], humans[:c]...)
	} else {
		return humans
	}
	return humans
}
