//問題文URL:https://paiza.jp/works/mondai/array_primer/array_primer__array_distance_step2

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
	A := num[1]
	B := num[2]
	m := map[string]int{}
	//N回座標が入力されるので格納していく、値は「x y」スペース区切りでx,yの値が入力される
	for i := 0; i < N; i++ {
		sc.Scan()
		strArr := strings.Split(sc.Text(), " ")
		m["x"+strconv.Itoa(i+1)], _ = strconv.Atoi(strArr[0])
		m["y"+strconv.Itoa(i+1)], _ = strconv.Atoi(strArr[1])
	}
	//マンハッタン距離で、A番目とB番目の座標間の距離を計算して出力する
	x := func() int {
		ans := m["x"+A] - m["x"+B]
		if ans < 0 {
			return ans * -1
		} else {
			return ans
		}
	}
	y := func() int {
		ans := m["y"+A] - m["y"+B]
		if ans < 0 {
			return ans * -1
		} else {
			return ans
		}
	}
	ans := x() + y()
	fmt.Println(ans)
}
