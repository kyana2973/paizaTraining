//問題文URL:https://paiza.jp/works/mondai/array_primer/array_primer__array_distance_step1

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
	N, _ := strconv.Atoi(sc.Text())
	m := map[string]int{}
	//固定値の座標x1,y1を格納する
	m["x1"] = 2
	m["y1"] = 3
	//N回座標が入力されるので格納していく、値は「x y」スペース区切りでx,yの値が入力される
	for i := 0; i < N; i++ {
		sc.Scan()
		strArr := strings.Split(sc.Text(), " ")
		//座標が連番になるようにする
		m["x"+strconv.Itoa(i+2)], _ = strconv.Atoi(strArr[0])
		m["y"+strconv.Itoa(i+2)], _ = strconv.Atoi(strArr[1])
	}
	//マンハッタン距離でx1,y1をベースにそれぞれ計算して改行区切りで出力する
	for i := 0; i < N; i++ {
		x := func() int {
			ans := m["x1"] - m["x"+strconv.Itoa(i+2)]
			if ans < 0 {
				return ans * -1
			} else {
				return ans
			}
		}
		y := func() int {
			ans := m["y1"] - m["y"+strconv.Itoa(i+2)]
			if ans < 0 {
				return ans * -1
			} else {
				return ans
			}
		}
		ans := x() + y()
		fmt.Println(ans)
	}
}
