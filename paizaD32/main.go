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
	keyWord := strings.Split(sc.Text(), " ")
	H, _ := strconv.Atoi(keyWord[0])
	W, _ := strconv.Atoi(keyWord[1])
	//AとBは、スペースを付加して9桁になるようにする
	A := fmt.Sprintf("%9s", keyWord[2])
	B := fmt.Sprintf("%9s", keyWord[3])

	//出力する値
	outputStr := "(" + A + ", " + B + ")"

	//行間の境界線
	var rowLine string
	for i := 0; i < (len(outputStr)+3)*W-3; i++ {
		rowLine = rowLine + "="
	}

	//列間に「 | 」を入れて、行間にも境界線を入れて値を出力する
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if j == W-1 {
				fmt.Println(outputStr)
			} else {
				fmt.Print(outputStr + " | ")
			}
		}
		if i != H-1 {
			fmt.Println(rowLine)
		}
	}
}
