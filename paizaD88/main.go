// 問題文URL:https://paiza.jp/works/mondai/array_utilization_primer/array_utilization_primer__kind

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())
	m := map[int]bool{}
	intArr := []int{}
	//スライスの中身で値が被っていない物は何種類あるかを出力する
	for i := 0; i < N; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		if !m[v] {
			m[v] = true
			intArr = append(intArr, v)
		}
	}
	fmt.Println(len(intArr))
}
