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
	//H = 出力する行数、W = 出力する列数、A = 括弧内の左に出力する値、B = 括弧内の右に出力する値
	HWAB := strings.Split(sc.Text(), " ")
	H, _ := strconv.Atoi(HWAB[0])
	W, _ := strconv.Atoi(HWAB[1])
	A := HWAB[2]
	B := HWAB[3]
	//各マスに出力する（A, B)を予め定義
	output := "(" + A + ", " + B + ")"
	//行と行の間に出力する境界線を定義する、「=」の数は上下出力される文字列の数に合わせる
	var rowLine string
	for i := 0; i < W*9-3; i++ {
		rowLine = rowLine + "="
	}

	//H、Wの数だけ出力し、行間の境界線として「=」、列間の境界線として「 | 」を出力する
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if (j + 1) == W {
				fmt.Println(output)
			} else {
				fmt.Print(output + " | ")
			}
		}
		if (i + 1) != H {
			fmt.Println(rowLine)
		}
	}
}
